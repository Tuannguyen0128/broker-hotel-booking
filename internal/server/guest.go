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

func (sv *server) GetGuests(ctx context.Context, req *proto.GetGuestsRequest) (*proto.GetGuestsResponse, error) {
	//result := &ClientResponse{}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("guests?id=%s=%s&citizenId=%s&phone=%s&page=%d&offset=%d", req.Id, req.CitizenId, req.Phone, req.Page, req.Offset)
	log.Println("Broker call to repo:", url)
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Content-Type", "application/json")

	// Do log
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
	var accounts *proto.GetGuestsResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&accounts)
	if err != nil {
		fmt.Println(err)
	}
	return accounts, nil
}

func (sv *server) CreateGuest(ctx context.Context, req *proto.Guest) (*proto.Guest, error) {
	//result := &ClientResponse{}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("guest")
	request, err := http.NewRequest("POST", url, nil)
	request.Header.Set("Content-Type", "application/json")

	// Do log
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
	var guest *proto.Guest
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&guest)
	if err != nil {
		fmt.Println(err)
	}
	return guest, nil
}

func (sv *server) DeleteGuest(ctx context.Context, req *proto.DeleteGuestRequest) (*proto.DeleteGuestResponse, error) {
	//result := &ClientResponse{}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("guest")
	request, err := http.NewRequest("DELETE", url, nil)
	request.Header.Set("Content-Type", "application/json")

	// Do log
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
	var response *proto.DeleteGuestResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println(err)
	}
	return response, nil
}

func (sv *server) UpdateGuest(ctx context.Context, req *proto.Guest) (*proto.Guest, error) {
	// Prepare log
	payload, _ := json.Marshal(req)
	jsonPayload := make([]byte, 0)
	if payload != nil {
		jsonPayload, _ = json.Marshal(payload)
	}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("guest")
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	request.Header.Set("Content-Type", "application/json")

	// Do log
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
	var guest *proto.Guest
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&guest)
	if err != nil {
		fmt.Println(err)
	}
	return guest, nil
}
