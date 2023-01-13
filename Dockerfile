FROM golang:1.18-alpine

WORKDIR /app

COPY dist .

EXPOSE 3000

RUN chmod +x dist

CMD [ "/app/dist" ]
