FROM golang:1.21 AS go-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./tmp/main ./cmd/main.go

EXPOSE 8080

CMD ["./tmp/main"]
