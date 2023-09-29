# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory to /app/server
WORKDIR /app/server

# Copy the Go application source code and dependencies
COPY ./server/ .

# Install Go dependencies (if using modules)
RUN go mod download

# Build the Go application
RUN go build -o main

# Make the main executable
RUN chmod +x main

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
