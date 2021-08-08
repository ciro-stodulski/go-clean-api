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
#### Versions

1.1.0

```
- refectory in module http
- ports for http 
- middleware structure for routes and global requests
- refectory in module db
- refectory part 2 in module http
- add test unit domain controller
- refectory interfaces in core 
- Add Docker and Docker composer
- add integration http 
- add test unit in integration http
```