FROM golang:1.21
WORKDIR /app/server
COPY ./server/ .
RUN go mod download
RUN go mod tidy
RUN go build -o main
RUN chmod +x main
EXPOSE 3000
CMD ["./main"]
