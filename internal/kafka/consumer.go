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
	batch := k.ConsumerConn.ReadBatch(1, 1e6) // fetch 10KB min, 1MB max
	b := make([]byte, 10e3)                   // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		//fmt.Println(string(b[:n]))
		ch <- b[:n]
	}
}
