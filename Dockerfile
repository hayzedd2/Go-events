# Use official Go image as base
FROM golang:1.21-alpine


# Set working directory
WORKDIR /app

# Install required system dependencies
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Create data directory for SQLite
RUN mkdir -p /app/data

# Set permissions for the data directory
RUN chmod 755 /app/data

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./main"]