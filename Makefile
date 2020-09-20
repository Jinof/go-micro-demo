
run_srv:
	go run user/cmd/srv/*.go

run_srv_broker_nats:
	go run user/cmd/srv/*.go --broker nats --broker_address :4222

run_api:
	go run user/cmd/api/*.go

run_api_broker_nats:
	go run user/cmd/api/*.go --broker nats --broker_address :4222



run_gateway:
	cd gateway && make run_without_casbin namespace=$(namespace) secret=$(secret)

run_gateway_without_casbin:
	cd gateway && make run_without_casbin namespace=$(namespace) secret=$(secret)

build_gateway:
	cd gateway && make build

.PHONY: run_srv run_api run_gateway run_gateway_without_casbin build_gateway