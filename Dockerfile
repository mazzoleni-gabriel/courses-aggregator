FROM golang:1.21-alpine

WORKDIR /app/src/courses-aggregator

ENV GOPATH=/app

COPY . /app/src/courses-aggregator

RUN go build cmd/api/main.go

ENTRYPOINT ["./main"]

EXPOSE 8080