FROM golang:1.17.7-alpine3.14

WORKDIR /go/src/app
COPY . .

RUN go get -d -v
RUN go build -v -o hats-for-parties

CMD ["./hats-for-parties"]
