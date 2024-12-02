# Use an official Go runtime as a parent image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy the current directory's content into the container at /app
COPY . .

# Download Go dependencies
RUN go mod init webhook-server && go mod tidy

# Compile the Go app
RUN go build -o webhook-server .

# Expose the port the app runs on
EXPOSE 8080

# Run the executable
CMD ["./webhook-server"]
