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

func (sv *server) GetBookings(ctx context.Context, req *proto.GetBookingsRequest) (*proto.GetBookingsResponse, error) {
	//result := &ClientResponse{}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("bookings?id=%s=%s&position=%s&page=%d&offset=%d", req.Id, req.Position, req.Page, req.Offset)
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
	var accounts *proto.GetBookingsResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&accounts)
	if err != nil {
		fmt.Println(err)
	}
	return accounts, nil
}

func (sv *server) CreateBooking(ctx context.Context, req *proto.Booking) (*proto.Booking, error) {
	//result := &ClientResponse{}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("booking")
	request, err := http.NewRequest("POST", url, nil)
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
	var booking *proto.Booking
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&booking)
	if err != nil {
		fmt.Println(err)
	}
	return booking, nil
}

func (sv *server) DeleteBooking(ctx context.Context, req *proto.DeleteBookingRequest) (*proto.DeleteBookingResponse, error) {
	//result := &ClientResponse{}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("booking/:id=%s", req.Id)
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
	var response *proto.DeleteBookingResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		fmt.Println(err)
	}
	return response, nil
}

func (sv *server) UpdateBooking(ctx context.Context, req *proto.Booking) (*proto.Booking, error) {
	// Prepare request
	payload, _ := json.Marshal(req)
	jsonPayload := make([]byte, 0)
	if payload != nil {
		jsonPayload, _ = json.Marshal(payload)
	}
	url := sv.CFG.RepoServer + api + fmt.Sprintf("booking")
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
	var booking *proto.Booking
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&booking)
	if err != nil {
		fmt.Println(err)
	}
	return booking, nil
}
