# go-architecture-api

*initial structure of an api that implements principles of a clean architecture in golang*

## Main stacks used

- http server (gin)
- http client (net/http)
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

### protoc command

[--proto_path=] path where is proto [src/presentation/grpc/services/user/find-user/proto,src/presentation/grpc/services/user/find-user/proto/find-user.proto]

[--go_out=] where proto buffer will be to create [plugins=grpc:src/presentation/grpc/services/user/find-user/pb]

exe: 

```bash
protoc --proto_path=src/presentation/grpc/services/user/find-user/proto src/presentation/grpc/services/user/find-user/proto/find-user.proto --go_out=plugins=grpc:src/presentation/grpc/services/user/find-user/pb
```

*program not found or is not executable*
try: 

Run vim ~/.bash_profile

export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin

Run source ~/.bash_profile


# evans 

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