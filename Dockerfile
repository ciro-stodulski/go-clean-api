FROM golang:1.18-alpine

WORKDIR /app

COPY dist .

EXPOSE 3000

CMD [ "/app/dist" ]
