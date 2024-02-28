package main

import (
	"broker-hotel-booking/configs"
	"broker-hotel-booking/internal/kafka"
	"broker-hotel-booking/internal/server"
)

func main() {
	// Init configs
	conf := configs.Load("./configs/config.yaml")
	kafkaClient := kafka.InitConnection(conf.KafkaServer, conf.KafkaBrokerTopic, conf.KafkaRepoTopic, 0)

	server.ListenAndServe("3001", kafkaClient, conf)
}
