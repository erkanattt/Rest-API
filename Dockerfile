FROM golang:1.24.2-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main cmd/main.go

FROM alpine:latest

COPY --from=builder /app/main /root/

CMD ["/root/main"]
