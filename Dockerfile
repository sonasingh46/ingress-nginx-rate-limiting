FROM golang:1.20

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
COPY cmd cmd

RUN CGO_ENABLED=0 go build cmd/main.go

EXPOSE 8080

CMD ["./main"]