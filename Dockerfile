# Build Env
FROM golang:1.13 AS build-env

ENV GO111MODULE=on

ADD . /go/src/github.com/damiannolan/eros

WORKDIR /go/src/github.com/damiannolan/eros

RUN go build -i -o app ./cmd/eros/

# Application Image
FROM gcr.io/distroless/base:latest

COPY --from=build-env /go/src/github.com/damiannolan/eros/app /usr/local/bin/app

CMD ["/usr/local/bin/app"]
