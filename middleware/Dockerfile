FROM golang:1.16

WORKDIR /go

ENV PORT=80
ENV BACKEND_URL=http://localhost:8080
EXPOSE $PORT

COPY go.mod src/go.mod
COPY main.go src/main.go
COPY config src/config
COPY search src/search
COPY autocomplete src/autocomplete

RUN cd src && go install .

CMD ["bin/myblog-search"]

