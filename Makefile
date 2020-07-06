.PHONY: run_srv
run_srv:
	go run user-srv/user/main.go

.PHONY: run_api
run_api:
	go run user-srv/api/main.go

.PHONY: run_gateway
run_gateway:
	cd gateway && go run main.go api --secret=$(secret)
