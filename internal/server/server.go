package server

import (
	"broker-hotel-booking/configs"
	"broker-hotel-booking/internal/kafka"
	"broker-hotel-booking/internal/proto"
	"broker-hotel-booking/internal/repositories"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	proto.AccountServiceServer
	kafkaClient *kafka.Kafka
	CFG         *configs.AppConfig
	repo        *repositories.Repositories
}

func NewSever(kafka *kafka.Kafka, appConfig *configs.AppConfig) *server {
	return &server{
		kafkaClient: kafka,
		CFG:         appConfig,
	}
}
func ListenAndServe(Port string, kafka *kafka.Kafka, config *configs.AppConfig) {
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
