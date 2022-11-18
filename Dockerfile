FROM golang:1.18-alpine

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

COPY ./ ./

RUN apk update && apk upgrade && apk add git
RUN go get github.com/cespare/reflex


RUN go mod download