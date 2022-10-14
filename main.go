package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fmt.Println("Go Websockets")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
