# Use Golang official image
FROM golang:1.22.4

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum from the src directory
COPY src/go.mod src/go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code from the src directory
COPY src/ ./

# Build the application
RUN go build -o src/main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["src/main"]
