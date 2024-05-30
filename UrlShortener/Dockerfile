FROM golang:alpine as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/natefinch/lumberjack
RUN go build -o main .
CMD ["/app/main"]
