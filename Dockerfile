FROM golang:1.17.2-alpine

WORKDIR /go/src/app
copy . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["fizzbuzzServer"]
