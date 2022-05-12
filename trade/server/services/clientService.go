package service

import (
	"trade/models"

	"github.com/gorilla/websocket"
)

type ClientService struct {
	Hub *models.Hub
}

func NewClientService() *ClientService {
	return &ClientService{
		Hub: &models.Hub{
			Broadcast:  make(chan []byte),
			Register:   make(chan *models.Client),
			Unregister: make(chan *models.Client),
			Clients:    make(map[*models.Client]bool),
		},
	}
}

func (s *ClientService) Register(conn *websocket.Conn) *models.Client {
	client := &models.Client{Hub: s.Hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client
	return client
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (s *ClientService) WritePump(c models.Client) {
	defer func() {
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

// func (s *ClientService) ReadPump(c models.Client, control controllers.Controller) {
// 	defer func() {
// 		c.Hub.Unregister <- &c
// 		c.Conn.Close()
// 	}()
// 	// c.conn.SetReadLimit(maxMessageSize)
// 	// c.conn.SetReadDeadline(time.Now().Add(pongWait))
// 	// c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
// 	for {
// 		_, message, err := c.Conn.ReadMessage()
// 		if err != nil {
// 			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
// 				log.Printf("error: %v", err)
// 			}
// 			break
// 		}
// 		event := models.Event{}
// 		err = json.Unmarshal([]byte(message), &event)
// 		if err != nil {
// 			log.Println("unmarshal write:", err)
// 			break
// 		}

// 		switch event.Name {
// 		case "order":
// 			control.OrderWs(c, event.Data)
// 		case "cancel":
// 			//FIXME:
// 		default:
// 		}
// 	}
// }

func (s *ClientService) Run() {
	for {
		select {
		case client := <-s.Hub.Register:
			s.Hub.Clients[client] = true
		case client := <-s.Hub.Unregister:
			if _, ok := s.Hub.Clients[client]; ok {
				delete(s.Hub.Clients, client)
				close(client.Send)
			}
		}
	}
}
