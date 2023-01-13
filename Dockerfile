FROM golang:1.18-alpine AS builder

# Export necessary port
EXPOSE 8082

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o go-hello .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/go-hello .

# Build a small image
FROM scratch

COPY --from=builder /dist/go-hello /

# Expose port 8082 to the outside world
EXPOSE 8082

# Command to run
ENTRYPOINT ["/go-hello"]

