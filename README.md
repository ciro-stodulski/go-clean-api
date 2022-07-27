# go-architecture-api

design based on clean architecture

https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

## Main stacks used

- http server (gin)
- http client (net/http)
- grpc server
- grpc client
- amqp server (rabbit mq/consumer)
- amqp client (rabbit mq/producer)
- mysql (gorm)
- redis 
- cron
- testify (unit tests)

## How init

- remove .sample of .env.sample
obs: you need to update the file with your env

## Started api

```bash
go run .
```

## Started api with container

```bash
docker-compose up
```

## Run tests
```bash
 go test ./...
```

# gRPC tips

### start server to test client grpc
```bash
cd server-client-grpc
go run .
```

### protoc command

How generate protobufjs

[--proto_path=] path where is proto [cmd/infra/integrations/grpc/notification/proto,cmd/infra/integrations/grpc/notification/proto/notification.proto]

[--go_out=] where proto buffer will be to create [plugins=grpc:cmd/infra/integrations/grpc/notification/pb]

infra layer exe: 

```bash
protoc --proto_path=cmd/infra/integrations/grpc/notification/proto cmd/infra/integrations/grpc/notification/proto/notification.proto --go_out=plugins=grpc:cmd/infra/integrations/grpc/notification/pb
```

interface layer exe: 

```bash
protoc --proto_path=cmd/interface/grpc/notification/proto cmd/interface/grpc/notification/proto/notification.proto --go_out=plugins=grpc:/home/santa-fe/Documents/playground/myDev/go-architecture-api/cmd/interface/grpc/notification/pb
```

*program not found or is not executable*
try: 

Run 

```bash
vim ~/.bash_profile
```

```bash
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin
```

Run

```bash
source ~/.bash_profile
```


## evans (test to server grpc)

```bash
evans -r --host localhost -p 50055
```

```bash
show service
```

```bash
service FindUserService
```

```bash
call FindUser

id (TYPE_STRING) => 1
{
  "user": {
    "name": "test",
    "email": "test"
  }
}
```



#### Current version

## [2.2.1]

```
- fix refactoring integrations folders
- add client to grpc 
- add service to host grpc
- update version go 1.18
- fix changelog and add loggers
- add new rote delete user
- fix connection redis
```
