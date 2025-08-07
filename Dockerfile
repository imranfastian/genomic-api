# Use Go image for runtime (development)
FROM golang:1.24.4-alpine

WORKDIR /app
RUN apk add --no-cache git

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .
EXPOSE 8080

CMD ["go", "run", "main.go"]
