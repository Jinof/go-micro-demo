# go-micro-demo

该 demo 展示如何通过 api 调用 grpc server

# Usage

Run the Custom gateway

    make run_gateway secret="Your secret"
    
Run the user-srv

    make run_srv
    
Run the api layer
 
    make run_api
    
# Call the service

    curl -H 'Content-Type: application/json' -H 'Authorization: passport Token' -d '{"name": "John"}' http://localhost:8080/user/call

    {
        "status": 0,
        "message": "成功调用User.Call",
        "data": "Hello johnaa"
    }