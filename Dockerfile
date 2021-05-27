FROM golang:1.16

WORKDIR /go/src/github.com/Bastorx
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]