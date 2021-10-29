FROM golang:1.14.1

WORKDIR /go/src/app
COPY . .

EXPOSE 3000