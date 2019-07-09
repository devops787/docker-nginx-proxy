FROM golang:1.12.6-alpine

MAINTAINER devops787

WORKDIR /go/src/docker-nginx-proxy
COPY . /go/src/docker-nginx-proxy

RUN apk --no-cache add git dep
RUN rm -rf ./vendor && dep ensure && go build -o app

EXPOSE 3000

ENTRYPOINT ["./app"]