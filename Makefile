proto-gen:
	protoc --proto_path=protobufs  --go-grpc_out=internal --go_out=internal protobufs/account.proto
	protoc --proto_path=protobufs  --go-grpc_out=internal --go_out=internal protobufs/staff.proto
	protoc --proto_path=protobufs  --go-grpc_out=internal --go_out=internal protobufs/booking.proto
	protoc --proto_path=protobufs  --go-grpc_out=internal --go_out=internal protobufs/guest.proto

run:
	go run main.go