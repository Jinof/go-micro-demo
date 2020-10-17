
run_srv:
	cd user/cmd/srv && make run

run_srv_broker_nats:
	cd user/cmd/srv && make run_broker_nats

run_api:
	cd user/cmd/api && make run

run_api_broker_nats:
	cd user/cmd/api && make run_broker_nats

run_gateway:
	cd gateway && make run_without_casbin namespace=$(namespace) secret=$(secret)

run_gateway_without_casbin:
	cd gateway && make run_without_casbin namespace=$(namespace) secret=$(secret)

build_gateway:
	cd gateway && make build

build_api:
	cd user/cmd/api && make build

build_srv:
	cd user/cmd/srv && make build

.PHONY: run_srv run_api run_gateway run_gateway_without_casbin build_gateway build_api build_srv