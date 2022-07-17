##### Stage 1 #####

### Use golang:1.15 as base image for building the application
FROM golang:1.17-alpine as builder

MAINTAINER Aceng rifik91@gmail.com

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod tidy

COPY . .

EXPOSE 3000

RUN go build

CMD ["./restful-cake-store"]