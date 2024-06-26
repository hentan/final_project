FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main ./cmd

EXPOSE 8080

CMD ["./main"]