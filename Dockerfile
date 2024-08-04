FROM golang:1.22.5 AS builder

# Set the working directory
WORKDIR /app/server

# Copy the go.mod and go.sum files first for better caching
COPY ./server/go.mod ./server/go.sum ./

RUN go mod download

# Copy the rest of the source code
COPY ./server/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a smaller base image for the final stage
FROM gcr.io/distroless/base-debian11 AS final

# Set the working directory in the new image
WORKDIR /app/server

# Copy the binary from the builder image
COPY --from=builder /app/server/main .

# Expose the port the app runs on
EXPOSE 3000

# Run the binary
CMD ["/app/server/main"]