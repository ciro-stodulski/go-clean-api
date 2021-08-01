FROM golang:1.16-alpine

WORKDIR /app
EXPOSE 8080

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build .

ENTRYPOINT ["./go-api"]