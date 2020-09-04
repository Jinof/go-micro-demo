module github.com/Jinof/go-micro-demo

go 1.14

require (
	github.com/casbin/casbin/v2 v2.8.1
	github.com/casbin/xorm-adapter/v2 v2.0.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/micro/v2 v2.9.3
	github.com/spf13/viper v1.6.3 // indirect
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
