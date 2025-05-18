# Build stage
FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

# Build the binary with correct name
RUN go build -o smtp2zoho main.go

# Final minimal image
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/smtp2zoho /app/smtp2zoho

# Define entrypoint
ENTRYPOINT ["/app/smtp2zoho"]
