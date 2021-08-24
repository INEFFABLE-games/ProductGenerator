FROM golang:latest

COPY . /user/ProductsGenerator
WORKDIR /user/ProductsGenerator
CMD go run ./main.go

EXPOSE 9092