package services

import "log"

func HandleMessage(msg string) (Resp, error) {
	log.Println("Received message:", msg)

	return Resp{
		raw: msg,
	}, nil
}
