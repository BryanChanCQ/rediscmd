FROM golang:1.23.0-alpine3.20 as builder
WORKDIR /app
COPY ./ /app
ENV GOPROXY='http://goproxy.cn'
RUN go build -o bin/redisCmd main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/bin/redisCmd /usr/local/bin

