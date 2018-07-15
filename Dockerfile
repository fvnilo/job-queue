FROM golang:1.10-alpine

RUN mkdir -p /go/src/github.com/nylo-andry/job-queue
WORKDIR /go/src/github.com/nylo-andry/job-queue

ADD . .

RUN go build -o main .