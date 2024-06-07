# Указываем базовый образ
FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod .

COPY . .

RUN go build -o main ./cmd/api/

WORKDIR /dist

RUN cp /app/main .

EXPOSE 8080

CMD ["/dist/main"]