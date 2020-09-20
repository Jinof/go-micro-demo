
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v2/api/proto

PB_PATH=./pb

.PHONY: proto
proto:
	rm -rf ${gen}/*.proto
	protoc --proto_path=${GOPATH}/src:${PB_PATH} \
	--micro_out=${MODIFY}:genproto \
	--go_out=${MODIFY}:genproto \
	api/user.proto
	protoc --proto_path=${GOPATH}/src:${PB_PATH} \
	--micro_out=${MODIFY}:genproto \
	--go_out=${MODIFY}:genproto \
	srv/user.proto
	protoc --proto_path=${PB_PATH} \
	--micro_out=${MODIFY}:genproto \
	--go_out=${MODIFY}:genproto \
	event/event.proto

.PHONY: build
build: proto
	go build -o user-service ./api/*.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-service:latest
