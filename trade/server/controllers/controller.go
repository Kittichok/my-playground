package controllers

import (
	"encoding/json"
	"net/http"
	"trade/models"
	service "trade/services"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

type Controller struct {
	serv       *service.Service
	clientServ *service.ClientService
}

func NewController(serv *service.Service, clientServ *service.ClientService) *Controller {
	return &Controller{
		serv:       serv,
		clientServ: clientServ,
	}
}

type OrderInput struct {
	OrderType string  `json:"orderType"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int64   `json:"quantity"`
}

func (c *Controller) OrderWs(client models.Client, data interface{}) ([]byte, error) {
	input := OrderInput{}
	mapstructure.Decode(data, &input)
	// if err != nil {
	// 	return nil, err
	// }
	if input.OrderType == "buy" {
		return c.serv.BuyOrder(client, input.Name, input.Price, input.Quantity)
	}
	return c.serv.SellOrder(client, input.Name, input.Price, input.Quantity)
}

func (c *Controller) Register(conn *websocket.Conn) *models.Client {
	return c.clientServ.Register(conn)
}

func (c *Controller) WritePump(client models.Client) {
	c.clientServ.WritePump(client)
}

type OrdersResp struct {
	Sell []models.Ticket `json:"sell"`
	Buy  []models.Ticket `json:"buy"`
}

func (c *Controller) GetOrders(w http.ResponseWriter, r *http.Request) {
	buy, sell := c.serv.GetOrders()
	orders := &OrdersResp{
		Sell: sell,
		Buy:  buy,
	}

	resp, _ := json.Marshal(orders)
	w.Write(resp)
}
