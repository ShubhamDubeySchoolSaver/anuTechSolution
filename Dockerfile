# Build the Go app
FROM golang:alpine as build

# Set the working directory
WORKDIR /app

# Copy all files from the current directory into the container
COPY . .

# Install Go modules required for the project
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/natefinch/lumberjack

# Build the Go binary
RUN go build -o main .

# Use a minimal base image to reduce size
FROM alpine:latest

WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=build /app/main .

# Copy the SQL script into the container (if needed for later execution)
COPY data.sql /docker-entrypoint-initdb.d/data.sql

# Create the log directory and set permissions
RUN mkdir -p /app/log && chmod -R 777 /app/log

# Add the database URL as environment variables
ENV MYSQL_HOST=autorack.proxy.rlwy.net
ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=FOmbtmYqkOmVAzEJXcBbIYfKScUHwAkr
ENV MYSQL_DATABASE=skool_saver
ENV MYSQL_PORT=27122

# Expose the application's port
EXPOSE 8080

# Start the application
CMD ["./main"]
