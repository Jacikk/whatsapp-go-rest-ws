# Use the official Golang image as the base image
FROM golang:alpine3.21
# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build entgo schemas
RUN go generate ./internals/server/ent

# Build the Go application
RUN go build -o entry cmd/main.go

# Expose the application port
EXPOSE 5000

# Command to run the application
CMD ["./entry"]