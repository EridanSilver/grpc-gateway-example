APP_NAME ?= grpc-gateway-example

all:
	mkdir -p "pkg"
	mkdir -p "pkg/pb"
	protoc -I/usr/local/include -I. \
		-I./vendor.proto \
		--grpc-gateway_out=logtostderr=true:./pkg/pb \
		--swagger_out=allow_merge=true,merge_file_name=api:. \
		--go_out=plugins=grpc:./api_pb ./*.proto