# Start by using an official Golang image as the build environment
FROM golang:1.20-alpine AS builder

LABEL stage="builder" version="1.0" description="asciiWeb: A Go-based web application that serves ASCII art text."

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go build -o asciiWeb .

# Use a smaller Alpine base image for the final stage
FROM alpine:latest

LABEL description="This is the final stage for the asciiWeb application using an Alpine base."

# Update package lists and install bash
RUN apk update && apk add bash

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/asciiWeb /app/asciiWeb

# Set the command to run the built binary
CMD ["./asciiWeb"]
