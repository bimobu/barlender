# Start from golang base image
FROM golang:1.23.6-bookworm

# Set working directory
WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Run the executable
CMD ["air"]