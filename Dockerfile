FROM golang:1.16

WORKDIR /go/src/github.com/Bastorx
ADD . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 4000

CMD ["app"]