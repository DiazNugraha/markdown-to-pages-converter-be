# Start from the official Golang base image
FROM golang:1.22.0

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

COPY .env .env

# Build the Go app
RUN go build -o main .

# Expose port 3000 to the outside world
EXPOSE ${PORT:-3000}

# Command to run the executable
CMD ["./main"]
