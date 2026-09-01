# ============================================
# STAGE 1: Build
# ============================================
# Use Go 1.22 on Alpine Linux to compile our app
FROM golang:1.22-alpine AS builder

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum (dependency files)
COPY go.mod go.sum ./

# Download dependencies
# This happens during build, not during container run
RUN go mod download

# Copy application source code
COPY main.go .

# Build the binary
# This creates a single executable file
RUN go build -o nigerian-food-api main.go

# ============================================
# STAGE 2: Runtime (lightweight)
# ============================================
# Use minimal Alpine Linux (no Go compiler)
FROM alpine:3.18

# Set working directory
WORKDIR /app

# Copy ONLY the compiled binary from Stage 1
# Not the Go compiler, not the source code
# Just the executable
COPY --from=builder /app/nigerian-food-api .

# Set environment variables
# These can be overridden at runtime with docker run -e
ENV PORT=8080
ENV ENVIRONMENT=production

# Document that app listens on port 8080
EXPOSE 8080

# Health check
# Docker will run this command every 30 seconds
# If it fails 3 times, mark container as unhealthy
HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

# Default command when container starts
CMD ["./nigerian-food-api"]
