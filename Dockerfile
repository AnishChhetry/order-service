FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o order-api main.go

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/order-api .
EXPOSE 4000
CMD ["./order-api"]