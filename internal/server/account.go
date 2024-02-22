package server

import (
	"broker-hotel-booking/internal/proto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

var (
	accountUrl = "/api/"
)

func (sv *server) GetAllAccount(ctx context.Context, req *proto.GetAccountsRequest) (*proto.GetAccountsResponse, error) {
	requestId, _ := uuid.NewUUID()
	message := &Message{
		Header: MessageHeader{
			ServiceName: "GetAllAccount",
			ReqID:       requestId.String(),
		},
		Body: req,
	}
	// Prepare request
	payload, _ := json.Marshal(message)
	jsonPayload := make([]byte, 0)
	if payload != nil {
		jsonPayload, _ = json.Marshal(payload)
	}

	//result := &ClientResponse{}
	url := sv.CFG.RepoServer + accountUrl + fmt.Sprintf("accounts?id=%s&staff_id=%s&username=%s&page=%d&offset=%d", req.ID, req.StaffID, req.Username, req.Page, req.Offset)
	log.Println("Broker call to repo:", url)
	request, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonPayload))
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
	var accounts *proto.GetAccountsResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&accounts)
	if err != nil {
		fmt.Println("Loi")
	}
	return accounts, nil
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
	// Prepare request
	payload, _ := json.Marshal(message)
	jsonPayload := make([]byte, 0)
	if payload != nil {
		jsonPayload, _ = json.Marshal(payload)
	}

	result := &ClientResponse{}
	url := ""
	request, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonPayload))
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
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result.RawData)
	account := &proto.Account{}
	byteRawdata := result.RawData.([]byte)
	json.Unmarshal(byteRawdata, account)
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
type ClientResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	RawData interface{} `json:"raw_data"`
}
