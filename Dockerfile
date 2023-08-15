# # Use an official Golang image as the base image
# FROM golang:alpine AS build

# # Set the working directory inside the container
# WORKDIR /app

# # Copy go.mod and go.sum to download and cache dependencies
# COPY go.mod go.sum ./

# # Download dependencies
# RUN go mod download

# # Copy the rest of the application code
# COPY . .

# # Build the application
# RUN go build 

# # Create a new stage for the final image
# FROM alpine:latest

# # Set the working directory in the final image
# WORKDIR /usr/local/bin

# # Copy the binary from the build stage to the final image
# COPY --from=build /app/twc .

# # Set execute permissions for the binary
# RUN chmod +x /usr/local/bin/twc

# # Expose the port as an environment variable
# ENV SERVER_PORT=8080

# # Set the entry point with the environment variable
# CMD ["sh", "-c", "/usr/local/bin/twc -port $SERVER_PORT"]
# Use the official Golang Alpine image as the base

FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o tictactoer

# Expose the port your application is listening on
EXPOSE 8080

# Set the entry point for the container to run the binary
CMD ["./tictactoer"]
