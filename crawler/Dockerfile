FROM curlimages/curl:7.77.0

USER root

RUN apk update
RUN apk upgrade
RUN apk add bash
RUN apk add libxslt
RUN apk add openjdk11

WORKDIR /crawler

ENV BACKEND_URL=http://host.docker.internal:8080

COPY bin /crawler/bin
COPY blog /crawler/blog

CMD ["bin/crawler.sh"]
