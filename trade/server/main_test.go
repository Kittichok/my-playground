package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/gorilla/websocket"
)

func TestMain(t *testing.T) {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	t.Run("should ordered", func(t *testing.T) {
		req := `{
			"name": "order",
			"data": {
				"name": "ONE",
				"price": 100,
				"quantity": 10, 
				"orderType": "sell"
			}
		}`

		err := connectAndSend(u, req)
		if err {
			return
		}
	})

	t.Run("should match", func(t *testing.T) {
		name := "BTX"
		req := fmt.Sprintf(`{
			"name": "order",
			"data": {
				"name": "%v",
				"price": 100,
				"quantity": 10, 
				"orderType": "sell"
			}
		}`, name)

		err := connectAndSend(u, req)
		if err {
			return
		}
		req = fmt.Sprintf(`{
			"name": "order",
			"data": {
				"name": "%v",
				"price": 100,
				"quantity": 10, 
				"orderType": "buy"
			}
		}`, name)

		err = connectAndSend(u, req)
		if err {
			return
		}
	})

	t.Run("get orders", func(t *testing.T) {
		resp, err := http.Get("http://localhost:8080/orders")
		if err != nil {
			t.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		r := string(body)
		log.Printf("orders: %s", r)
		t.Log(r)
	})
}

func connectAndSend(u url.URL, req string) bool {
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// done := make(chan struct{})

	// go func() {
	// 	defer close(done)
	// 	for {
	// 		_, message, err := c.ReadMessage()
	// 		if err != nil {
	// 			return
	// 		}
	// 		log.Printf("recv: %s", message)
	// 		done <- struct{}{}
	// 	}
	// }()

	// ticker := time.NewTicker(time.Second)
	// defer ticker.Stop()

	err = c.WriteMessage(websocket.TextMessage, []byte(req))
	if err != nil {
		log.Println("write:", err)
		return true
	}
	return true

	// for {
	// 	select {
	// 	case <-done:
	// 		c.Close()
	// 		return true
	// 	case t := <-ticker.C:
	// 		log.Println("write:", t)
	// 		err := c.WriteMessage(websocket.TextMessage, []byte(req))
	// 		if err != nil {
	// 			log.Println("write:", err)
	// 			return true
	// 		}
	// 	}
	// }
}
