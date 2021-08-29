# Build
FROM golang:1.16 AS build

WORKDIR /go

COPY go.mod src/go.mod
COPY go.sum src/go.sum
COPY main.go src/main.go
COPY convert src/convert

RUN cd /go/src && go build -o ../bin/convert-blog

# Deploy
FROM curlimages/curl:7.77.0

USER root

RUN apk update
RUN apk upgrade
RUN apk add bash
RUN apk add libxml2-utils
RUN apk add openjdk11

WORKDIR /crawler

ENV BACKEND_URL=http://host.docker.internal:8080

COPY bin /crawler/bin
COPY --from=build /go/bin/convert-blog /crawler/bin/convert-blog

CMD ["bin/crawler.sh"]
