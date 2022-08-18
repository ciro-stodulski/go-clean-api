FROM golang:alpine3.15

WORKDIR /app
EXPOSE 8080

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build .

ENTRYPOINT ["./go-api"]