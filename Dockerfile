FROM golang:1.22.5 AS builder

# Set the working directory
WORKDIR /app/server

# Copy the go.mod and go.sum files first for better caching
COPY ./server/go.mod ./server/go.sum ./

RUN go mod download

# Copy the rest of the source code
COPY ./server/ .
RUN go build -o main .
RUN chmod +x main

# Create a new, minimal image to run the binary
FROM ubuntu:22.04

# Set the working directory in the new image
WORKDIR /app/server

# Copy the binary from the builder image
COPY --from=builder /app/server/main .

EXPOSE 3000
CMD ["./main"]
