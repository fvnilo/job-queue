FROM golang:1.10-alpine as builder

RUN apk add --update make
RUN mkdir -p /go/src/github.com/nylo-andry/jobqueue
WORKDIR /go/src/github.com/nylo-andry/jobqueue

ADD . .

RUN make build

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /go/src/github.com/nylo-andry/jobqueue/bin/publisher .
COPY --from=builder /go/src/github.com/nylo-andry/jobqueue/bin/listener .