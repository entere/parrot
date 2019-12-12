module github.com/entere/parrot

go 1.13

replace github.com/entere/parrot => /Users/entere/github/parrot

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	go.uber.org/zap v1.12.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
