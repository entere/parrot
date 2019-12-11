# Auth 服务

登录认证服务，用于登录和退出登录

生成方式

```
micro new auth-web --namespace=com.island.code --alias=auth --type=web --gopath=false
```

## 入门指南

- [配置](#配置)
- [依赖](#依赖)
- [使用](#使用)

## 配置

- FQDN: com.island.code.web.auth
- Type: web
- Alias: auth

## 依赖

依赖 etcd

```
# install etcd
brew install etcd

# run etcd
etcd
```

## 使用

为了方便起见，还包括了Makefile

编译

```
make build
```

运行
```
./auth-web
```

使用Docker构建
```
make docker
```

## 直接调用服务测试

```

curl localhost:8088/auth/login -X POST -d 'loginName=judy&loginPwd=456'

curl localhost:8088/auth/login -X POST -d '{"loginName": "judy","loginPwd":"456"}' --header "Content-Type: application/json"

```

## 启动micro api网关调用服务

注意：先启动micro网关，然后启动 srv 和 web 服务，先后顺序不能错

1、启动micro网关

```
micro --registry=etcd --api_namespace=com.island.code.web  api --handler=web
```

2、启动 srv 服务

```
./auth-srv
```

3、启动 web 服务

```
./auth-web
```

4、curl 调用

```
curl localhost:8080/auth/login -X POST -d 'loginName=judy&loginPwd=456'
# curl --request POST   --url http://localhost:8080/auth/login   --header 'Content-Type: application/x-www-form-urlencoded'  --data 'loginName=judy&loginPwd=456'
```

