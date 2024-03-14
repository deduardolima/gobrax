FROM golang:1.21.1 as builder

WORKDIR /app

COPY go.mod go.sum ./

COPY cmd/server/.env ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

FROM scratch

COPY --from=builder /app/main .

CMD ["./main"]
