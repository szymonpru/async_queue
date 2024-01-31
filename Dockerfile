FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy only the files in the src/ directory
COPY src/ .

# Build the Go application
RUN go build -o app .
