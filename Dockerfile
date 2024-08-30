FROM golang:1.22.5 AS builder

WORKDIR /app/server
COPY ./server/go.mod ./server/go.sum ./

RUN go mod download


COPY ./server/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a smaller base image for the final stage
FROM gcr.io/distroless/base-debian11 AS final
WORKDIR /app/server
COPY --from=builder /app/server/main .
EXPOSE 3000
CMD ["/app/server/main"]