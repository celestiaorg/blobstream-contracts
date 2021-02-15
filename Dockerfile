#install packages for build layer
FROM golang:1.15-alpine as builder
RUN apk add --no-cache git gcc make perl jq libc-dev linux-headers

#build binary
WORKDIR /src
ENV GO111MODULE=on
COPY . .
RUN cd orchestrator && go mod download

#install binary
RUN cd orchestrator && make install

#build main container
FROM alpine:latest
RUN apk add --update --no-cache ca-certificates
RUN apk add curl
COPY --from=builder /go/bin/* /usr/local/bin/

#configure container
VOLUME /apps/data
WORKDIR /apps/data

#default command
CMD cd /root/.injectived/peggo/ && peggo_orchestrator
