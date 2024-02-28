package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func InitConnection(host string, producerTopic string, consumerTopic string, partition int) *Kafka {
	consumer, err := kafka.DialLeader(context.Background(), "tcp", host, consumerTopic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	producer, err := kafka.DialLeader(context.Background(), "tcp", host, producerTopic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	return &Kafka{
		ProducerConn: producer,
		ConsumerConn: consumer,
	}
}

func (k *Kafka) ReadMessage(ch chan []byte) {
	k.ConsumerConn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch, err := k.ConsumerConn.ReadMessage(1e6) // fetch 10KB min, 1MB max
	if err != nil {
		log.Println(err.Error())
	}
	ch <- batch.Value
}
