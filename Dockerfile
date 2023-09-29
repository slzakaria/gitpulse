# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory to /app/server
WORKDIR /app/server

# Copy the Go application source code and dependencies
COPY ./server/ .

# Build the Go application
RUN go build -o main

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
