package server

import (
	"broker-hotel-booking/internal/proto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	api = "/api/"
)

func (sv *server) GetStaffs(ctx context.Context, req *proto.GetStaffsRequest) (*proto.GetStaffsResponse, error) {
	//result := &ClientResponse{}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("staffs?id=%s=%s&position=%s&page=%d&offset=%d", req.Id, req.Position, req.Page, req.Offset)
	log.Println("Broker call to repo:", url)
	request, err := http.NewRequest("GET", url, nil)
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
	var accounts *proto.GetStaffsResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&accounts)
	if err != nil {
		fmt.Println(err)
	}
	return accounts, nil
}

func (sv *server) CreateStaff(ctx context.Context, req *proto.Staff) (*proto.Staff, error) {
	// Prepare request
	payload, _ := json.Marshal(req)
	jsonPayload := make([]byte, 0)
	if payload != nil {
		jsonPayload, _ = json.Marshal(payload)
	}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("staff")
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
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
	var staff *proto.Staff
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&staff)
	if err != nil {
		fmt.Println(err)
	}
	return staff, nil
}

func (sv *server) DeleteStaff(ctx context.Context, req *proto.DeleteStaffRequest) (*proto.DeleteStaffResponse, error) {
	// Prepare request
	url := sv.CFG.RepoServer + api + fmt.Sprintf("staff/:id=%s", req.Id)
	request, err := http.NewRequest("DELETE", url, nil)
	request.Header.Set("Content-Type", "application/json")

	// Do request
	client := http.Client{}
	log.Println("Broker call to repo:", request)
	resp, err := client.Do(request)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("Failed unable to reach the server.", err, url)
		return nil, err
	}

	// Response
	var response *proto.DeleteStaffResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println(err)
	}
	return response, nil
}

func (sv *server) UpdateStaff(ctx context.Context, req *proto.Staff) (*proto.Staff, error) {
	// Prepare request
	payload, _ := json.Marshal(req)
	jsonPayload := make([]byte, 0)
	if payload != nil {
		jsonPayload, _ = json.Marshal(payload)
	}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("staff")
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	request.Header.Set("Content-Type", "application/json")

	// Do request
	client := http.Client{}
	log.Println("Broker call to repo:", request)
	resp, err := client.Do(request)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		log.Println("Failed unable to reach the server.", err, request)
		return nil, err
	}

	// Response
	var staff *proto.Staff
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&staff)
	if err != nil {
		fmt.Println(err)
	}
	return staff, nil
}
