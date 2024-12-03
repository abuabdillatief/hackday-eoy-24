package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hackday/services"
)

var upgrader = websocket.Upgrader{}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}

		resp, err := services.HandleMessage(string(message))
		if err != nil {
			log.Println("Error during message handling:", err)
		}

		err = conn.WriteMessage(messageType, resp.Bytes())
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/socket", socketHandler)
	log.Println("websocket server started on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
