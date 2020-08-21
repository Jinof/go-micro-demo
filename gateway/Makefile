NAME=gateway

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
	go run ./*.go \
	--auth_namespace=${namespace} \
	api \
	--secret=$(secret) \
	--namespace=${namespace}

run_without_casbin:
	go run ./*.go \
	--auth_namespace=${namespace} \
	api \
	--secret=$(secret) \
	--namespace=${namespace} \
	--casbin_init_skip=true


.PHONY: build docker push run run_without_casbin