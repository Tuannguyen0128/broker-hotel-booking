package server

import (
	"broker-hotel-booking/internal/models"
	"broker-hotel-booking/internal/proto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

type KafkaRequest struct {
	ServiceName string      `json:"service_name"`
	Payload     interface{} `json:"payload"`
}

func (sv *server) GetAccounts(ctx context.Context, req *proto.GetAccountsRequest) (*proto.GetAccountsResponse, error) {
	// Prepare request
	url := sv.CFG.RepoServer + api + fmt.Sprintf("accounts?id=%s&staff_id=%s&username=%s&page=%d&offset=%d", req.Id, req.StaffId, req.Username, req.Page, req.Offset)

	kafkaRequest := KafkaRequest{
		ServiceName: "GetAccounts",
		Payload:     &req,
	}

	message, err := json.Marshal(kafkaRequest)
	if err != nil {
		return nil, err
	}

	// Do request
	sv.kafkaClient.SendMessage(message)
	if err != nil {
		log.Println("Failed unable to reach the server.", err, url)
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
		fmt.Println(err.Error())
	}
	if response.Error.Code != "" {
		return nil, errors.New(response.Error.Error())
	}
	if response.Body == nil {
		return &proto.GetAccountsResponse{Accounts: nil}, nil
	}
	log.Println(response)
	jsonbody, err := json.Marshal(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	var accounts = &proto.GetAccountsResponse{}
	if err := json.Unmarshal(jsonbody, &accounts.Accounts); err != nil {
		// do error check
		fmt.Println(err)
	}
	return accounts, nil
}
func (sv *server) CreateAccount(ctx context.Context, req *proto.Account) (*proto.CreateAccountResponse, error) {
	// Prepare request
	requestAcc := &models.Account{
		StaffId:    req.StaffId,
		Username:   req.Username,
		Password:   req.Password,
		UserRoleId: req.UserRoleId,
	}
	payload, err := json.Marshal(requestAcc)
	//jsonPayload := make([]byte, 0)
	if err != nil {
		return nil, err
	}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("account")
	log.Println("Broker call to repo:", url)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/json")

	// Do request
	client := http.Client{}
	resp, err := client.Do(request)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("Failed unable to reach the server.", err, url)
		return nil, err
	}

	// Response
	var response *models.Response
	respDecoder := json.NewDecoder(resp.Body)
	err = respDecoder.Decode(&response)
	if err != nil {
		fmt.Println(err)
	}
	if response.Error.Code != "" {
		return &proto.CreateAccountResponse{Id: ""}, errors.New(response.Error.Error())
	}
	if response.Body == nil {
		return &proto.CreateAccountResponse{}, nil
	}
	log.Println(response)
	jsonbody, err := json.Marshal(response.Body)
	if err != nil {
		// do error check
		fmt.Println(err)
	}
	var result *proto.CreateAccountResponse
	if err := json.Unmarshal(jsonbody, &result); err != nil {
		// do error check
		fmt.Println(err)
	}
	return result, nil
}

func (sv *server) UpdateAccount(ctx context.Context, req *proto.Account) (*proto.Account, error) {
	// Prepare request
	payload, _ := json.Marshal(req)
	jsonPayload := make([]byte, 0)
	if payload != nil {
		jsonPayload, _ = json.Marshal(payload)
	}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("account")
	request, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonPayload))
	request.Header.Set("Content-Type", "application/json")

	// Do request
	client := http.Client{}
	log.Println("Broker call to repo:", request, url)
	resp, err := client.Do(request)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("Failed unable to reach the server.", err, url)
		return nil, err
	}

	// Response
	var account *proto.Account
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&account)
	if err != nil {
		fmt.Println(err)
	}
	return account, nil
}

func (sv *server) DeleteAccount(ctx context.Context, req *proto.DeleteAccountRequest) (*proto.DeleteAccountResponse, error) {
	// Prepare request
	payload, _ := json.Marshal(req)
	jsonPayload := make([]byte, 0)
	if payload != nil {
		jsonPayload, _ = json.Marshal(payload)
	}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("account")
	request, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonPayload))
	request.Header.Set("Content-Type", "application/json")

	// Do request
	client := http.Client{}
	log.Println("Broker call to repo:", request, url)
	resp, err := client.Do(request)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("Failed unable to reach the server.", err, url)
		return nil, err
	}

	// Response
	var result *proto.DeleteAccountResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result, nil
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
