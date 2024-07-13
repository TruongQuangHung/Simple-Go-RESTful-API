# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine

RUN apk add --no-cache \
    # Important: required for go-sqlite3
    gcc \
    # Required for Alpine
    musl-dev

WORKDIR /app

COPY . .

RUN go mod download

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build  -o /out/main ./

# Expose port 8080 for incoming traffic
EXPOSE 8080

# Define the command to run the app when the container starts
ENTRYPOINT ["/out/main"]