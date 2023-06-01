# Start from the latest golang base image as the builder
FROM golang:latest AS builder

# Set the Current Working Directory inside the builder
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the working directory inside the builder
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Now create the final runtime image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the built binary from the builder
COPY --from=builder /app/main .

# This container is meant to be interactive, so we use an ENTRYPOINT to allow parameters to be passed easily.
ENTRYPOINT [ "./main" ]
