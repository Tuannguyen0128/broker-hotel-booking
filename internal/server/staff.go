package server

import (
	"broker-hotel-booking/internal/models"
	"broker-hotel-booking/internal/proto"
	"broker-hotel-booking/internal/repositories"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func (sv *server) GetStaffs(ctx context.Context, req *proto.GetStaffsRequest) (*proto.GetStaffsResponse, error) {
	// Prepare request

	kafkaRequest := KafkaRequest{
		ServiceName: "GetStaffs",
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
		return &proto.GetStaffsResponse{Staffs: nil}, nil
	}
	if response.Error.Code != "" {
		return nil, errors.New(response.Error.Error())
	}
	if response.Body == nil {
		return &proto.GetStaffsResponse{Staffs: nil}, nil
	}

	jsonBody, err := json.Marshal(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Store json detail to mongodb
	data := models.LogData{
		ID:             primitive.ObjectID{},
		RequestDetail:  string(message),
		ResponseDetail: string(jsonBody),
		CreatedDate:    time.Now(),
	}

	repo := repositories.New()
	repo.LogRepo.InsertRequest(&data)
	if err != nil {
		log.Println("Failed to store log to database.", err)
	}

	// Decode response body
	var staffs = &proto.GetStaffsResponse{}
	if err := json.Unmarshal(jsonBody, &staffs.Staffs); err != nil {
		// do error check
		fmt.Println(err)
	}
	return staffs, nil
}

//func (sv *server) CreateStaff(ctx context.Context, req *proto.Staff) (*proto.CreateStaffResponse, error) {
//	// Prepare request
//	payload, _ := json.Marshal(req)
//	jsonPayload := make([]byte, 0)
//	if payload != nil {
//		jsonPayload, _ = json.Marshal(payload)
//	}
//	url := sv.CFG.RepoServer + api + fmt.Sprintf("staff")
//	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
//	request.Header.Set("Content-Type", "application/json")
//
//	// Do request
//	client := http.Client{}
//	log.Println("Broker call to repo:", request, url)
//	resp, err := client.Do(request)
//	if resp != nil && resp.Body != nil {
//		defer resp.Body.Close()
//	}
//	if err != nil {
//		log.Println("Failed unable to reach the server.", err, url)
//		return nil, err
//	}
//
//	// Response
//	var staff *proto.CreateStaffResponse
//	decoder := json.NewDecoder(resp.Body)
//	err = decoder.Decode(&staff)
//	if err != nil {
//		fmt.Println(err)
//	}
//	return staff, nil
//}
//
//func (sv *server) DeleteStaff(ctx context.Context, req *proto.DeleteStaffRequest) (*proto.DeleteStaffResponse, error) {
//	// Prepare request
//	url := sv.CFG.RepoServer + api + fmt.Sprintf("staff/:id=%s", req.Id)
//	request, err := http.NewRequest("DELETE", url, nil)
//	request.Header.Set("Content-Type", "application/json")
//
//	// Do request
//	client := http.Client{}
//	log.Println("Broker call to repo:", request)
//	resp, err := client.Do(request)
//	if resp != nil && resp.Body != nil {
//		defer resp.Body.Close()
//	}
//	if err != nil {
//		log.Println("Failed unable to reach the server.", err, url)
//		return nil, err
//	}
//
//	// Response
//	var response *proto.DeleteStaffResponse
//	decoder := json.NewDecoder(resp.Body)
//	err = decoder.Decode(&response)
//	if err != nil {
//		fmt.Println(err)
//	}
//	return response, nil
//}
//
//func (sv *server) UpdateStaff(ctx context.Context, req *proto.Staff) (*proto.Staff, error) {
//	// Prepare request
//	payload, _ := json.Marshal(req)
//	jsonPayload := make([]byte, 0)
//	if payload != nil {
//		jsonPayload, _ = json.Marshal(payload)
//	}
//	url := sv.CFG.RepoServer + api + fmt.Sprintf("staff")
//	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
//	request.Header.Set("Content-Type", "application/json")
//
//	// Do request
//	client := http.Client{}
//	log.Println("Broker call to repo:", request)
//	resp, err := client.Do(request)
//	if resp != nil && resp.Body != nil {
//		defer resp.Body.Close()
//	}
//	if err != nil {
//		log.Println("Failed unable to reach the server.", err, request)
//		return nil, err
//	}
//
//	// Response
//	var staff *proto.Staff
//	decoder := json.NewDecoder(resp.Body)
//	err = decoder.Decode(&staff)
//	if err != nil {
//		fmt.Println(err)
//	}
//	return staff, nil
//}
