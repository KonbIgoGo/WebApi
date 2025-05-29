FROM golang:1.24.3 AS build-stage

WORKDIR /usr/WebApi

COPY . .