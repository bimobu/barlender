# Start from golang base image
FROM golang:1.23.6-bookworm

# Set working directory
WORKDIR /app

RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Run the executable
CMD ["air"]