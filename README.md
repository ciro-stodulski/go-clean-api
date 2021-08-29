# go-architecture-api

##### *initial structure of an api that implements principles of a clean architecture in golang*

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
#### Current version

2.0.0

```
- refectory interfaces in core
- add work for manager cronjobs
- add cronjob
- add test for cron job list user
- add cache structure
- refectory test unit list user
- Added unit testing in the cache at the infra layer
- Fix cache in use case list users
- Added structure init for consumer 
- Added create user and list user with amqp presentation 
- Added tests unit
- Added envs for rabbit mq, redis, integration json place holder
- Added structure init for producer  
```