
run_srv:
	go run user/cmd/srv/main.go

run_api:
	go run user/cmd/api/main.go

run_gateway:
	cd gateway && make run_without_casbin namespace=$(namespace) secret=$(secret)

run_gateway_without_casbin:
	cd gateway && make run_without_casbin namespace=$(namespace) secret=$(secret)

build_gateway:
	cd gateway && make build

.PHONY: run_srv run_api run_gateway run_gateway_without_casbin build_gateway