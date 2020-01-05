FROM golang:1.13.5-alpine3.11

WORKDIR ./src
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["main.go"]
