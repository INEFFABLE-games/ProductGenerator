FROM golang:latest

COPY . /user/AliExpress/ProductGenerator
WORKDIR /user/AliExpress/ProductGenerator
CMD go run ./main.go

EXPOSE 9092