package main

import (
	"broker-hotel-booking/internal/kafka"
	"broker-hotel-booking/internal/server"
	"fmt"
)

func main() {
	kafkaClient := kafka.InitConnection("localhost:29092", "broker-repo", "broker-repo", 0)
	kafkaClient.SendMessage([]byte("Send message to kafka"))
	kafkaClient.SendMessage([]byte("Send message to kafka 2"))
	ch := make(chan []byte)
	go kafkaClient.ReadMessage(ch)
	fmt.Println(string(<-ch))
	fmt.Println(string(<-ch))
	server.ListenAndServe("3001", kafkaClient)
}
