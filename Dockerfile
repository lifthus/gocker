# syntax=docker/dockerfile:1

FROM golang:1.21.3-bullseye
WORKDIR /app
COPY . .
RUN go build -o gocker cmd/gocker/main.go
CMD ["./gocker"]
EXPOSE 9595