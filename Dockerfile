# Build stage
FROM golang:1.22-alpine as build
LABEL maintainer="app"
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main cmd/main.go

# Final stage
FROM alpine:latest
WORKDIR /
COPY --from=build /app/main /main
COPY .env .env
CMD ["./main"]