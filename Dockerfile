FROM golang:1.22-alpine AS build
WORKDIR /app
COPY go.mod main.go ./
RUN go build -o probe .
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/probe ./
EXPOSE 8080
CMD ["./probe"]
