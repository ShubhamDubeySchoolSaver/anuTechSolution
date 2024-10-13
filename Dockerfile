#Using offical maven image as a parent image
FROM golang:alpine as build

#Setting the working directory to /app
WORKDIR /app
MAINTAINER Vinayak
#Copy the current directory contents into the container at current directory
COPY . .
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/natefinch/lumberjack
#Install the mvn command for maven 
RUN go build -o main .