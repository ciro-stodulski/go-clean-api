# FROM golang:1.18-alpine

# WORKDIR /app

# EXPOSE 3000

# COPY dist .

# CMD ./dist


# WORKDIR /usr/src/dist

# # pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

# COPY . .
# RUN ls
# RUN go build -v -o dist /usr/local/bin ./...

# CMD ["./app"]


FROM golang:1.18-alpine as builder

# Set Environment Variables
ENV HOME /app
ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build app
RUN go build -a -installsuffix cgo -o dist .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage

EXPOSE 3000

CMD [ "./main" ]