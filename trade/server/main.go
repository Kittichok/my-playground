package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	controller "trade/controllers"
	"trade/models"
	"trade/repository"
	service "trade/services"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

var control = &controller.Controller{}

func serveWs(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	client := control.Register(c)
	go control.WritePump(*client)
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		event := models.Event{}
		err = json.Unmarshal([]byte(message), &event)
		if err != nil {
			log.Println("unmarshal write:", err)
			break
		}
		_, err = process(event, client)
		if err != nil {
			log.Println("process write:", err)
			break
		}
	}
}

func process(event models.Event, client *models.Client) ([]byte, error) {
	switch event.Name {
	case "order":
		return control.OrderWs(*client, event.Data)
	case "cancel":
		//FIXME:
		return nil, nil
	default:
		return nil, nil
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	repoSellTicket := repository.NewMemRepository[models.Ticket]()
	repoBuyTicket := repository.NewMemRepository[models.Ticket]()
	serv := service.NewService(&repoSellTicket, &repoBuyTicket)
	clientServ := service.NewClientService()
	go clientServ.Run()
	control = controller.NewController(serv, clientServ)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(w, r)
	})
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		control.GetOrders(w, r)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))

}
