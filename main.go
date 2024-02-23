package main

import (
	"broker-hotel-booking/config"
	"broker-hotel-booking/internal/server"
)

func main() {

	// Init config
	conf := config.Load("./config/config.yaml")
	server.ListenAndServe("3001", nil, conf)
}
