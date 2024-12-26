# Start from the official Go image
FROM golang:1.22.4

# Set the working directory inside the container
WORKDIR /

# Copy the source code into the container
COPY go.mod go.sum ./

RUN go mod download


COPY . .

# Clean

  # Install dependencies and build the application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
