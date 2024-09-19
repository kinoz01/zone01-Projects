# Start by using an official Golang image as the build environment
FROM golang:1.20-alpine AS builder

LABEL stage="builder" version="1.0" description="groupie"

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

# Build the Go application
RUN go build -o groupie .

# Use a smaller Alpine base image for the final stage
FROM alpine:latest

# Update package lists and install bash
RUN apk update && apk add bash

ENV PORT 8080
ENV APIPORT 4000

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/groupie /app/groupie
COPY --from=builder /app/frontend /app/frontend
COPY --from=builder  /app/API/locations.json /app/API/locations.json
COPY --from=builder  /app/API/apidata.json /app/API/apidata.json
COPY --from=builder  /app/server/apiLinks.json /app/server/apiLinks.json

# Set the command to run the built binary
CMD ["./groupie"]

# How to run
# curl -L https://fly.io/install.sh | sh
# export FLYCTL_INSTALL="/home/kino/.fly"
# export PATH="$FLYCTL_INSTALL/bin:$PATH"
# . ~/.bashrc
# flyctl deploy
# Visit your newly deployed app at https://groupie-tracker.fly.dev/
