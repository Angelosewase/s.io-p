package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("New connection established")

		so.On("disconnect", func() {
			log.Println("Connection closed")
		})

		so.On("message", func(msg string) {
			log.Println("Received message:", msg)
			so.Emit("message", msg)
		})
	})

	r.GET("/socket.io/*any", gin.WrapH(server))

	r.Run(":8080")
}