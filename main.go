package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {return true },
}

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("error handling socket %d",err )
			return
		}
		defer conn.Close()

		// Handle WebSocket messages
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				break
			}
			log.Printf("Received message: %s", message)
			err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, client!"))
			if err != nil {
				log.Println(err)
				break
			}
		}
	})

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}