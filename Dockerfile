FROM golang:1.22.4 AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./cmd/ -o stress_test

FROM scratch
WORKDIR /app
COPY --from=builder /app/cmd/stress_test ./
ENTRYPOINT ["./stress_test"]
