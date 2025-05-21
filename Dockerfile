FROM golang:1.24.3

WORKDIR /usr/WebApi/

COPY . .
RUN go mod tidy