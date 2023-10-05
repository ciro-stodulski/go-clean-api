# go-clean-api

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
- server sent event

## How init

- remove .sample of .env.sample
obs: you need to update the file with your env

## Started 

```bash
docker-compose up -d
```

```bash
goose -dir migrations up
```

```bash
go run .
```

obs: need install goose to run migrations

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

## Run tests
```bash
 go test ./...
```

## gRPC tips

### start server to test client grpc
```bash
cd server-client-grpc
go run .
```

### protoc command

How generate protobufjs

[--proto_path=] path where is proto [cmd/infra/integration/grpc/notification/proto,cmd/infra/integration/grpc/notification/proto/notification.proto]

[--go_out=] where proto buffer will be to create [plugins=grpc:cmd/infra/integration/grpc/notification/pb]

infra layer exe: 

```bash
protoc --proto_path=cmd/infra/integration/grpc/notification/proto cmd/infra/integration/grpc/notification/proto/notification.proto --go_out=plugins=grpc:cmd/infra/integration/grpc/notification/pb
```

presetation layer exe: 

```bash
protoc --proto_path=cmd/presentation/grpc/notification/proto cmd/presentation/grpc/notification/proto/test-notification.proto --go_out=plugins=grpc:/home/santa-fe/Documents/playground/myDev/go-architecture-api/cmd/presentation/grpc/notification/pb
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


## Command line

- list commands 

```bash
go run ./cmd/main/modules/cli/main  list-commands
```

- run command <command>

```bash
go run ./cmd/main/modules/cli/main run-command  list-users
```

- help

```bash
go run ./cmd/main/modules/cli/main -h
```

*obs: -h, --help   help for this command*


### Current version


## [3.2.1]

- refactor application errors
- refactor plural names
- refactor http module (fix validation request body with gin and add swagger package)
- refactor partial domain layer
- add swagger docs for rest api
- refactor injection dependency presentation layer
- interface with generic
- server sent event
