FROM golang:1.8

RUN mkdir /app

ADD ./src /app

WORKDIR /app

RUN go get -u github.com/gorilla/mux

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]
