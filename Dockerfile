# Use the official Golang image as the base image
FROM golang:1.8

# Set the Current Working Directory inside the container
WORKDIR /Users/nattapat/Desktop/tln

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
