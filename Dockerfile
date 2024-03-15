FROM golang:1.20-buster as builder

WORKDIR /app

COPY go.mod go.sum ./
COPY cmd/server/.env /app/.env

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

FROM debian:buster-slim

WORKDIR /root/

COPY --from=builder /app/main .
COPY cmd/server/.env ./

EXPOSE 8080

CMD ["./main"]
