FROM golang:1.19-alpine

WORKDIR /app

COPY dist .

EXPOSE 3000

CMD [ "/dist" ]


# WORKDIR /usr/src/dist

# # pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# COPY go.mod go.sum ./
# RUN go mod download && go mod verify

# COPY . .
# RUN ls
# RUN go build -v -o dist /usr/local/bin ./...

# CMD ["./app"]


# FROM golang:1.18-alpine

# # Set Environment Variables
# ENV HOME /app
# ENV CGO_ENABLED 0
# ENV GOOS linux

# WORKDIR /app

# COPY dist .

# WORKDIR /root/

# EXPOSE 3000

# CMD [ "./dist" ]