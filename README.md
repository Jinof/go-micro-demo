# go-micro-demo

该 demo 展示如何通过 api 调用 grpc server

# Usage

Run the Micro Api

    micro api
    
Run the user-srv

    go run user-srv/main.go user-srv/plugin.go
    
Run the gateway
 
    go run gateway/main.go
    
# Call the service

    curl -H 'Content-Type: application/json' -d '{"name": "John"}' http://localhost:8080/user/call

    {"message": "Hello John"}