# Use Golang base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the app will run on (e.g., 8080)
EXPOSE 8080

# Start the Go application
CMD ["./main"]
