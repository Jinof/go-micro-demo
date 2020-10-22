NAME=user_api

IMAGE_TAG=${tag}
IMAGE_NAME=jinof/go-micro-demo-$(NAME)

build:
	CGO_ENABLED=0 go build -installsuffix cgo -o $(NAME) *.go

docker:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) .
	docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(IMAGE_NAME):latest

push:
	docker push $(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(IMAGE_NAME):latest

run:
	go run ./*.go

run_broker_nats:
	go run *.go --broker nats --broker_address Jinof:JinofGo@127.0.0.1:4222

.PHONY: build docker push run run_without_casbin