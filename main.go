package main

import (
	"golang.org/x/net/websocket"
	"net/http"
)

type Message struct {
	Message   string `json:"message"`
	TimeStamp int    `json:"timestamp"`
}

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	msg := &Message{"Hello from the other side", 9999}
	websocket.JSON.Send(ws, msg)
}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
