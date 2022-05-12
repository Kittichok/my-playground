package models

import (
	"time"

	"github.com/gorilla/websocket"
)

type Ticket struct {
	ID        int64     `json:"-"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Quantity  int64     `json:"quantity"`
	Client    Client    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type Client struct {
	Hub *Hub

	Conn *websocket.Conn

	Send chan []byte
}

type Hub struct {
	Clients map[*Client]bool

	Broadcast chan []byte

	Register chan *Client

	Unregister chan *Client
}

type Event struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}
