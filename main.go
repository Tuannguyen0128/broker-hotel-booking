package main

import (
	"broker-hotel-booking/config"
	"broker-hotel-booking/internal/kafka"
	"broker-hotel-booking/internal/server"
	"fmt"
)

func main() {
	// Init config
	conf := config.Load("./config/config.yaml")
	kafkaClient := kafka.InitConnection(conf.KafkaServer, conf.KafkaBrokerTopic, conf.KafkaRepoTopic, 0)
	fmt.Println("Sent to repo")
	kafkaClient.SendMessage([]byte("Broker send message to repo"))

	server.ListenAndServe("3001", kafkaClient, conf)
}
