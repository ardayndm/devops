FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod main.go main_test.go ./

RUN go mod download

RUN CGO_ENABLED=0 go build -o myapp main.go

FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache curl

COPY --from=builder /app/myapp .


CMD ["./myapp"]

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 CMD curl -f http://localhost:8080/health || exit 1