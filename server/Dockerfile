FROM golang:1.12

RUN mkdir /app

ADD ./src /app

WORKDIR /app

RUN go get -u github.com/gorilla/mux

RUN go get -u github.com/go-redis/redis


RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]
