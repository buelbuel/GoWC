# Use the official Go image as the build environment
FROM golang:1.22 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Copy the example config to the actual config
RUN cp config.example.toml config.toml

# Build the Go app
RUN go build -o main ./cmd/main.go

# Use a minimal image for the final build
FROM debian:bookworm-slim

# Set the working directory in the final image
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main /app/main

# Copy the config file
COPY --from=builder /app/config.toml /app/config.toml

# Copy the resources directory
COPY --from=builder /app/resources /app/resources

# Expose port 4001 to the outside world
EXPOSE 4000

# Command to run the application
CMD ["./main"]
