FROM golang:1.18-alpine

WORKDIR /app
EXPOSE 3000

COPY dist .
#COPY go-clean-api .

ENTRYPOINT ["./dist"]