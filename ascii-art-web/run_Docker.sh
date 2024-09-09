#!/bin/bash

# Colors for styling
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print success message
print_success() {
    echo -e "${GREEN}$1${NC}"
}

# Function to print error message
print_error() {
    echo -e "${RED}$1${NC}"
}

# Function to print info message
print_info() {
    echo -e "${BLUE}$1${NC}"
}

# Function to print a warning message
print_warning() {
    echo -e "${YELLOW}$1${NC}"
}

# Prompt for image and container names
read -p "Enter image name: " IMAGE_NAME
read -p "Enter container name: " CONTAINER_NAME


# Function to validate the port number
validate_port() {
    if [[ $1 =~ ^[0-9]+$ ]] && [ "$1" -ge 2000 ] && [ "$1" -le 65535 ]; then
        return 0
    else
        return 1
    fi
}

# Prompt for port number and validate it
while true; do
    read -p "Enter port number (2000-65535): " PORT
    if validate_port "$PORT"; then
        break
    else
        print_error "Invalid port number. Please enter a number between 2000 and 65535."
    fi
done


# Build the Docker image
print_info "Building Docker image..."
if docker build -t "$IMAGE_NAME" .; then
    print_success "Docker image built successfully!"
else
    print_error "Failed to build Docker image."
    exit 1
fi

# Check if a container with the same name already exists and remove it
EXISTING_CONTAINER=$(docker ps -aq -f name="$CONTAINER_NAME")

if [ -n "$EXISTING_CONTAINER" ]; then
    print_warning "Removing existing container with name '$CONTAINER_NAME'..."
    docker rm -f "$EXISTING_CONTAINER"
fi

# Run the Docker container
print_info "Running Docker container on port $PORT..."
# -p <host_port>:<container_port>
if docker run -d -p "$PORT:$PORT" --name "$CONTAINER_NAME" -e PORT="$PORT" "$IMAGE_NAME"; then
    print_success "Docker container is running!"
    print_info "Starting server at: http://127.0.0.1:$PORT"
else
    print_error "Failed to start Docker container."
    exit 1
fi

# Check logs of the container
print_info "Fetching container logs..."
docker logs "$CONTAINER_NAME"

# Ask the user if they want to enter the container's bash shell
read -p "Do you want to enter the container's bash shell? (y/n): " ENTER_SHELL

if [[ "$ENTER_SHELL" =~ ^[Yy]$ ]]; then
    docker exec -it "$CONTAINER_NAME" /bin/bash
else
    print_warning "You can enter the container using: docker exec -it $CONTAINER_NAME /bin/bash"
fi

# Additional command to stop the container if needed
print_warning "To stop the container, run: docker stop $CONTAINER_NAME"

# Clean up command
print_warning "To remove all the containers and images run: ./rm_Docker.sh"

# Remove intermediate image
print_warning 'To remove only intermediate images run: docker rmi golang:1.20-alpine $(docker images --filter "label=stage=builder" -q)'
