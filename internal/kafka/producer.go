package kafka

import (
	"github.com/segmentio/kafka-go"
	"log"
)

func (k *Kafka) SendMessage(message []byte) {
	_, err := k.ProducerConn.WriteMessages(
		kafka.Message{Value: message},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
}

type Kafka struct {
	ProducerConn *kafka.Conn
	ConsumerConn *kafka.Conn
}
