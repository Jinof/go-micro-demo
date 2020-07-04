
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v2/api/proto

PB_PATH=./pb
GEN_PATH=github.com/go-micro-demo/user-srv/${gen}

.PHONY: proto
proto:
	rm -rf ${gen}/*.proto
	protoc --proto_path=${GOPATH}/src:${PB_PATH} \
	--micro_out=Msrv/greeter.proto=${GEN_PATH}/srv,Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v2/api/proto:${gen} \
	--go_out=Msrv/greeter.proto=${GEN_PATH}/srv,Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v2/api/proto:${gen} \
	--validate_out="lang=go:${gen}" \
	api/user.proto
	protoc --proto_path=${GOPATH}/src:${PB_PATH} \
	--micro_out=${gen} \
	--go_out=:${gen} \
	--validate_out="lang=go:${gen}" \
	srv/user.proto
    

.PHONY: build
build: proto
	go build -o user-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-service:latest
