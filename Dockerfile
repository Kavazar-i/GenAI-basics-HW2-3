FROM golang:1.21-alpine AS builder

WORKDIR /build

# Install necessary build tools
RUN apk add --no-cache gcc musl-dev

# Copy the source code
COPY . .

# Download all dependencies and generate go.sum
RUN go mod download && \
    go mod tidy && \
    go mod verify

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Create the final image
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /build/main .

# Create a non-root user
RUN adduser -D -g '' appuser
USER appuser

# Command to run the executable
CMD ["./main"] 