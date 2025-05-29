FROM golang:1.24.3 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o webapi ./cmd/webapi

FROM alpine:latest AS app
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=build  /src/webapi .
EXPOSE 8080
ENTRYPOINT [ "./webapi" ]