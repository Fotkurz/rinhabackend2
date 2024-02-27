# Use the official golang image as a base image
FROM golang:1.21.5

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY ./go.mod .
COPY ./go.sum .

# Download and install dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

