# Auth 服务

登录认证服务，用于生成jwt,redis 存储jwt，删除jwt

生成方式

```
micro new auth-srv --namespace=com.island.code --alias=auth --type=srv --gopath=false
```

## 入门指南

- [配置](#配置)
- [依赖](#依赖)
- [使用](#使用)

## 配置

- FQDN: com.island.code.srv.auth
- Type: srv
- Alias: auth

## 依赖

依赖etcd、redis、mysql

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
./auth-srv
```

使用Docker构建
```
make docker
```

## 调用服务

1、micro call 调用

```
micro --registry etcd  call com.island.code.srv.auth Auth.MakeAccessToken '{"userID":"543987654"}'
micro --registry etcd  call com.island.code.srv.auth Auth.DelUserAccessToken '{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Nzg1NTg1OTQsImp0aSI6IjU0Mzk4NzY1NCIsImlhdCI6MTU3NTk2NjU5NCwiaXNzIjoiY29tLmlzbGFuZC5jb2RlIiwibmJmIjoxNTc1OTY2NTk0LCJzdWIiOiI1NDM5ODc2NTQifQ.jcxGnEFbIcSG48FU-WG72X-hE7aYsNNl83TGI6F716A"}'
micro --registry etcd  call com.island.code.srv.auth Auth.QueryUserByName '{"loginName":"entere"}'

```

2、curl 直接调用服务

通过/rpc这个固定url可以绕过rpc处理器直接对服务进行访问，例如


