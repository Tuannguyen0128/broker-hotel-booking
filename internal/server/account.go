package server

import (
	"broker-hotel-booking/internal/proto"
	"context"
	"encoding/json"
	"github.com/google/uuid"
)

func (sv *server) GetAllAccount(ctx context.Context, req *proto.GetAllAccountRequest) (*proto.GetAllAccountResponse, error) {
	requestId, _ := uuid.NewUUID()
	message := &Message{
		Header: MessageHeader{
			ServiceName: "GetAllAccount",
			ReqID:       requestId.String(),
		},
		Body: req,
	}
	bytes, _ := json.Marshal(message)
	sv.kafkaClient.SendMessage(bytes)
	reponse := make(chan []byte)
	sv.kafkaClient.ReadMessage(reponse)
	var accounts []*proto.Account
	json.Unmarshal(<-reponse, accounts)
	res := &proto.GetAllAccountResponse{
		Accounts: accounts,
	}
	return res, nil
}
func (sv *server) GetAccountByCitizenID(ctx context.Context, req *proto.GetAccountByCitizenIDRequest) (*proto.Account, error) {
	requestId, _ := uuid.NewUUID()
	message := &Message{
		Header: MessageHeader{
			ServiceName: "GetAccountByCitizenID",
			ReqID:       requestId.String(),
		},
		Body: req,
	}
	response := make(chan []byte)
	bytes, _ := json.Marshal(message)
	sv.kafkaClient.SendMessage(bytes)
	sv.kafkaClient.ReadMessage(response)
	var account *proto.Account
	json.Unmarshal(<-response, account)
	return account, nil
}

type MessageHeader struct {
	ServiceName string `json:"service_name"`
	ReqID       string `json:"req_id"`
}

type Message struct {
	Header MessageHeader
	Body   interface{}
}
type Accounts struct {
	list []Account
}
