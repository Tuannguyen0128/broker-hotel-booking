package main

import "broker-hotel-booking/internal/server"

func main() {
	server.ListenAndServe("3001")
}
