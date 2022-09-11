FROM golang:1.16-alpine

WORKDIR /app
EXPOSE 3000

COPY dist .
#COPY go-clean-api .

ENTRYPOINT ["./dist"]