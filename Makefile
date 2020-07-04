.PHONY: run_srv
run_srv:
	go run user-srv/user/main.go

.PHONY: run_api
run_api:
	go run user-srv/api/main.go

.PHONY: run_gateway
run_gateway:
	go run gateway/main.go api  --secret="Your Secret"
