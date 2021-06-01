FROM golang:1.16

ARG BUILD_ENV=Production

WORKDIR /go/src/github.com/Bastorx/test_lbc
ADD . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o .

EXPOSE 4000

RUN echo ${BUILD_ENV}
RUN if [ "${BUILD_ENV}" = "Development" ]; \
    then curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air; else go build -o .;fi

CMD ["test_lbc"]