FROM golang:1.18

WORKDIR /app/

COPY / /app

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
