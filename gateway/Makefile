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

.PHONY: build docker push