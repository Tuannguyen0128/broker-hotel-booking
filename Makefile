proto-gen:
	protoc --proto_path=protobufs  --go-grpc_out=internal --go_out=internal protobufs/account.proto

run:
	go run main.go