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
