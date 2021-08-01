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

1.0.0 - init project

```
- add entity init structure
- add entity test
- add use case init structure 
- add use case test
- add structure module init server http 
- add controller init structure
- add container init structure
- add inject controller in container
- add inject use case in container
- add repository init structure 
- add database module init structure 
- add repository test
```