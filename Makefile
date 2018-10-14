protoc:
	protoc -I. --go_out=plugins=grpc:. data.proto