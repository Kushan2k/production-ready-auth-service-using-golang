FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go mod tidy

COPY . .


RUN go build -o main ./cmd
RUN swag init -g cmd/main.go

FROM alpine:3.19
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
ENTRYPOINT ["./main"]
