package service

import (
	"log"
	"trade/models"
	"trade/repository"
)

type Service struct {
	repoSellTicket *repository.Repository[models.Ticket]
	repoBuyTicket  *repository.Repository[models.Ticket]
}

func NewService(repoSellTicket *repository.Repository[models.Ticket],
	repoBuyTicket *repository.Repository[models.Ticket]) *Service {
	return &Service{
		repoSellTicket: repoSellTicket,
		repoBuyTicket:  repoBuyTicket,
	}
}

func (s *Service) BuyOrder(client models.Client, name string, price float64, quantity int64) ([]byte, error) {
	log.Print("ordering:", name, price)
	// find a match
	// if found, create a trade
	// if not found, create a new order
	for i, t := range s.repoSellTicket.DB {
		if t.Name == name && t.Price == price {
			if (t.Quantity - quantity) < 0 {
				quantity -= t.Quantity
				continue
			}
			s.repoSellTicket.DB = RemoveIndex(s.repoSellTicket.DB, i)
			msg := []byte("You buy have a match!")
			t.Client.Send <- msg
			client.Send <- msg
			quantity -= t.Quantity
			if quantity > 0 {
				continue
			}
			break
		}
	}

	if quantity > 0 {
		s.repoBuyTicket.DB = append(s.repoBuyTicket.DB, models.Ticket{
			Name:     name,
			Price:    price,
			Quantity: quantity,
			Client:   client,
		})
		msg := []byte("You buy order recv!")
		client.Send <- msg
	}

	return nil, nil
}

func (s *Service) SellOrder(client models.Client, name string, price float64, quantity int64) ([]byte, error) {
	for i, t := range s.repoBuyTicket.DB {
		if t.Name == name && t.Price == price {
			if (t.Quantity - quantity) < 0 {
				quantity -= t.Quantity
				continue
			}
			s.repoBuyTicket.DB = RemoveIndex(s.repoBuyTicket.DB, i)
			msg := []byte("You sell have a match!")
			t.Client.Send <- msg
			client.Send <- msg
			// s.Notify.SendToClient(clientID, msg)
			quantity -= t.Quantity
			if quantity > 0 {
				continue
			}
			break
		}
	}

	if quantity > 0 {
		s.repoSellTicket.DB = append(s.repoSellTicket.DB, models.Ticket{
			Name:     name,
			Price:    price,
			Quantity: quantity,
			Client:   client,
		})
		msg := []byte("You sell order recv!")
		client.Send <- msg
	}

	return nil, nil
}

func (s *Service) GetOrders() ([]models.Ticket, []models.Ticket) {
	return s.repoBuyTicket.DB, s.repoSellTicket.DB
}

func RemoveIndex(s []models.Ticket, index int) []models.Ticket {
	return append(s[:index], s[index+1:]...)
}
