# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the files
COPY . .

# Build your binary
ENV CGO_ENABLED=0
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o cloud-strife-user main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/cloud-strife-user .

# Copy config and env files if needed
COPY config.json .

# Expose the port your app uses
EXPOSE 5001

# Run the binary
CMD ["./cloud-strife-user"]