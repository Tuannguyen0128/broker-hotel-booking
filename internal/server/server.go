package server

import (
	"broker-hotel-booking/config"
	"broker-hotel-booking/internal/kafka"
	"broker-hotel-booking/internal/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type server struct {
	proto.AccountServiceServer
	kafkaClient *kafka.Kafka
	CFG         *config.AppConfig
}

func NewSever(kafka *kafka.Kafka, appConfig *config.AppConfig) *server {
	return &server{
		kafkaClient: kafka,
		CFG:         appConfig,
	}
}
func ListenAndServe(Port string, kafka *kafka.Kafka, config *config.AppConfig) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", Port))
	if err != nil {
		log.Fatalln(err)
	}

	// Init grpc
	s := grpc.NewServer()
	proto.RegisterAccountServiceServer(s, NewSever(kafka, config))
	fmt.Println("Server connecting...")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalln("error while serve", err)

	}
}

type Account struct {
	ID           uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Fullname     string    `gorm:"size:20;not null" json:"fullname"`
	Email        string    `gorm:"size:50;not null;unique" json:"email"`
	Password     string    `gorm:"size:60;not null" json:"password"`
	CreatedAt    time.Time `gorm:"autoCreatedTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdatedTime" json:"updated_at"`
	MerchantCode string    `json:"merchantcode"`
}
