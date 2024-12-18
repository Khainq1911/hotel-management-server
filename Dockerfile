# Start with a lightweight Golang image
FROM golang:alpine as builder

# Set environment variables to speed up builds
ENV CGO_ENABLED=0 GOOS=linux

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY .env ./
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application binary
RUN go build -o main

# Use a smaller base image for the final stage
FROM alpine:3.18

# Set the working directory in the final image
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 3000

# Command to run the application
CMD ["./main"]
