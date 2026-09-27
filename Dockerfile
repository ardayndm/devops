FROM golang:1.26 AS builder

WORKDIR /app

COPY go.mod main.go main_test.go ./

RUN go mod download

RUN go build -o myapp main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/myapp .

CMD ["./myapp"]