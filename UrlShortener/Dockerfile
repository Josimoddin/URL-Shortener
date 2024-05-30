FROM golang:alpine as builder
RUN mkdir /masterapp
ADD . /masterapp
WORKDIR /masterapp
RUN go get github.com/natefinch/lumberjack
RUN go build -o main .
CMD ["/app/main"]
