FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the current directory contents into the container
COPY . .

# Install necessary dependencies
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/natefinch/lumberjack

# Build the Go application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]