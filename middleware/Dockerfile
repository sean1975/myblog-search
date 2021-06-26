FROM golang:1.16

WORKDIR /go

ENV PORT=80
ENV BACKEND_URL=http://localhost:8080/search
EXPOSE $PORT

COPY src/main.go src/main.go

RUN go build -o bin/main src/main.go

CMD ["bin/main"]

