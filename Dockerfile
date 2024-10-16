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

# Install MySQL client to run SQL script
RUN apk add --no-cache mysql-client

# Copy the SQL script into the container
COPY data.sql /docker-entrypoint-initdb.d/data.sql

# Add the database URL as an environment variable (customize as needed)
ENV MYSQL_HOST=mysql.railway.internal
ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=FOmbtmYqkOmVAzEJXcBbIYfKScUHwAkr
ENV MYSQL_DATABASE=skool_saver
ENV MYSQL_PORT=3306

# Wait for MySQL to be ready and then run the SQL script
CMD sh -c "sleep 10 && mysql -h $MYSQL_HOST -u $MYSQL_USER -p$MYSQL_PASSWORD $MYSQL_DATABASE < /docker-entrypoint-initdb.d/data.sql && ./main"
