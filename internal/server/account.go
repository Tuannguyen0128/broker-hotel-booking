package server

import (
	"broker-hotel-booking/internal/models"
	"broker-hotel-booking/internal/proto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"time"
)

type KafkaRequest struct {
	ServiceName string      `json:"service_name"`
	Payload     interface{} `json:"payload"`
}

func (sv *server) GetAccounts(ctx context.Context, req *proto.GetAccountsRequest) (*proto.GetAccountsResponse, error) {
	// Prepare request

	kafkaRequest := KafkaRequest{
		ServiceName: "GetAccounts",
		Payload:     &req,
	}

	message, err := json.Marshal(kafkaRequest)
	if err != nil {
		return nil, err
	}
	log.Println("Request", string(message))

	// Do request
	sv.kafkaClient.SendMessage(message)
	if err != nil {
		log.Println("Failed unable to reach the server.", err)
		return nil, err
	}

	// Response
	ch := make(chan []byte)
	go sv.kafkaClient.ReadMessage(ch)
	responseByte := <-ch
	fmt.Println(string(responseByte))
	var response = &models.Response{}
	err = json.Unmarshal(responseByte, response)
	if err != nil {
		return &proto.GetAccountsResponse{Accounts: nil}, nil
	}
	if response.Error.Code != "" {
		return nil, errors.New(response.Error.Error())
	}
	if response.Body == nil {
		return &proto.GetAccountsResponse{Accounts: nil}, nil
	}
	// Decode response body
	jsonBody, err := json.Marshal(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var accounts = &proto.GetAccountsResponse{}
	if err := json.Unmarshal(jsonBody, &accounts.Accounts); err != nil {
		// do error check
		fmt.Println(err)
	}
	return accounts, nil
}
func (sv *server) CreateAccount(ctx context.Context, req *proto.Account) (*proto.CreateAccountResponse, error) {
	// Prepare request
	log.Println("Request", req)
	requestAcc := &models.Account{
		StaffId:    req.StaffId,
		Username:   req.Username,
		Password:   req.Password,
		UserRoleId: req.UserRoleId,
	}
	kafkaRequest := KafkaRequest{
		ServiceName: "CreateAccount",
		Payload:     requestAcc,
	}

	message, err := json.Marshal(kafkaRequest)
	if err != nil {
		return nil, errors.New("Unable to marshal request")
	}

	// Do request
	sv.kafkaClient.SendMessage(message)

	// Response
	ch := make(chan []byte)
	go sv.kafkaClient.ReadMessage(ch)
	responseByte := <-ch
	fmt.Println(string(responseByte))
	var response = &models.Response{}
	err = json.Unmarshal(responseByte, response)
	if err != nil {
		return nil, errors.New("Unable to decode response")
	}

	if response.Error.Code != "" {
		return nil, errors.New(response.Error.Error())
	}
	if response.Body == nil {
		return &proto.CreateAccountResponse{Id: ""}, nil
	}

	// Decode response body
	jsonBody, err := json.Marshal(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var account = &proto.CreateAccountResponse{}
	if err := json.Unmarshal(jsonBody, &account); err != nil {
		// do error check
		fmt.Println(err)
	}
	return account, nil
}

func (sv *server) UpdateAccount(ctx context.Context, req *proto.Account) (*proto.Account, error) {
	// Prepare request
	log.Println("Request", req)
	requestAcc := &models.Account{
		Id:         req.Id,
		StaffId:    req.StaffId,
		Username:   req.Username,
		Password:   req.Password,
		UserRoleId: req.UserRoleId,
		UpdatedAt:  time.Now().String(),
	}
	kafkaRequest := KafkaRequest{
		ServiceName: "UpdateAccount",
		Payload:     requestAcc,
	}

	message, err := json.Marshal(kafkaRequest)
	if err != nil {
		return nil, errors.New("Unable to marshal request")
	}

	// Do request
	sv.kafkaClient.SendMessage(message)

	// Response
	ch := make(chan []byte)
	go sv.kafkaClient.ReadMessage(ch)
	responseByte := <-ch
	fmt.Println(string(responseByte))
	var response = &models.Response{}
	err = json.Unmarshal(responseByte, response)
	if err != nil {
		return nil, errors.New("Unable to decode response")
	}

	if response.Error.Code != "" {
		return nil, errors.New(response.Error.Error())
	}
	if response.Body == nil {
		return &proto.Account{}, nil
	}

	// Decode response body
	jsonBody, err := json.Marshal(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var account = &proto.Account{}
	if err := json.Unmarshal(jsonBody, &account); err != nil {
		// do error check
		fmt.Println(err)
	}
	return account, nil
}

func (sv *server) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*proto.DeleteAccountResponse, error) {
	// Prepare request
	log.Println("Request", req)
	requestAcc := &proto.DeleteAccountRequest{
		Id: req.Id,
	}
	kafkaRequest := KafkaRequest{
		ServiceName: "DeleteAccount",
		Payload:     requestAcc,
	}

	message, err := json.Marshal(kafkaRequest)
	if err != nil {
		return nil, errors.New("Unable to marshal request")
	}

	// Do request
	sv.kafkaClient.SendMessage(message)

	// Response
	ch := make(chan []byte)
	go sv.kafkaClient.ReadMessage(ch)
	responseByte := <-ch
	fmt.Println(string(responseByte))
	var response = &models.Response{}
	err = json.Unmarshal(responseByte, response)

	if err != nil {
		return nil, errors.New("Unable to decode response")
	}
	if response.Error.Code != "" {
		return nil, errors.New(response.Error.Error())
	}
	if response.Body == nil {
		return &proto.DeleteAccountResponse{}, nil
	}

	// Decode response body
	jsonBody, err := json.Marshal(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var account = &proto.DeleteAccountResponse{}
	if err := json.Unmarshal(jsonBody, &account); err != nil {
		// do error check
		fmt.Println(err)
	}
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

type ClientResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	RawData interface{} `json:"raw_data"`
}
