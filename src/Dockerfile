# Using a golang base image
FROM golang:alpine as builder

#Adding src directory contents to /go/src/app
ADD ./src /go/src/app

#Switching the work directory to /go/src/app
WORKDIR /go/src/app

ENV PORT=3001
CMD ["go", "run", "main.go"]
