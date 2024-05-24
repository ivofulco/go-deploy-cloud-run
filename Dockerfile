# Build stage
FROM golang:1.22.3-alpine as builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags '-s -w' \
    -o cloudrun cmd/main.go

# Run stage
FROM alpine as production
COPY --from=builder /app/cloudrun .
ENTRYPOINT ["./cloudrun"]
