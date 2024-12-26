# Start from the official Go image
FROM golang:1.22.4

# Set the working directory inside the container
WORKDIR /

# Copy the source code into the container
COPY . .

# Clean

  # Install dependencies and build the application
RUN go mod download
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
