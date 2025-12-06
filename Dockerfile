
##############################
# Build stage
##############################
FROM golang:alpine AS builder

# Enable modules and static build
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /app

# Copy go module files first (better layer caching)
COPY go.mod  ./ 

# Download dependencies (if go.mod exists)
RUN go mod download

# Copy source
COPY main.go .

# Build the binary
RUN go build -o server .

##############################
# Runtime stage
##############################
FROM scratch

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

# App listens on 8080 by default
ENV PORT=8080
EXPOSE 8080

# Run the app
ENTRYPOINT ["./server"]

