# go-micro-demo

下面是micro官方已经写的一些优秀的例子

| [start-kit](https://github.com/micro-in-cn/starter-kit) 
| [x-gateway](https://github.com/micro-in-cn/x-gateway)
 
当然你也可以直接看[awesome-micro](https://github.com/micro/awesome-micro)
 
架构
    
    req -> gateway -> api -> srv

srv为rpc服务, 通过api层抽象出接口, 再用gateway转发请求.    

优点:
1. gateway 实际为 micro api, 具有micro原生的服务发现的能力, 可自行调用 api.
1. 服务扩展时只需拓展api和srv层, gateway无需更改.
    
gateway的auth通过api plugin实现, 并使用了 casbin 做鉴权.
由于micro.Auth还未完成,所以只能自己写auth.
注意: 
    
    Header中的Authorization必须被重写为 'Authorization: Bearer xxx' 的形式.
    看micro初始化时的源码, 你可以看到micro.Auth会被使用,它要求Authorization必须
    为以上形式, xxx 可以为任何值, 因为不起作用(当micro.Auth完成后推荐使用它, 不用
    自己写token相关逻辑了😀).

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