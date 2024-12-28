FROM golang:1.23-alpine

# Install required dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /src/app

# Install air (live reload tool)
RUN go install github.com/air-verse/air@latest

# Copy project files
COPY . .

# Download Go dependencies
RUN go mod tidy
