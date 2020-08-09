.PHONY: run_srv
run_srv:
	go run user-srv/user/main.go

.PHONY: run_api
run_api:
	go run user-srv/api/main.go

.PHONY: run_gateway
run_gateway:
	cd gateway && go run main.go --auth_namespace=${namespace} api --secret=$(secret) --namespace=${namespace}

.PHONY: build
build_gateway:
	cd gateway && make build