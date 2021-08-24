package main

import (
	"ProductsGenerator/src/server"
	"log"
)

func main(){
	err := server.StartServer();if err != nil{
		log.Fatal(err)
	}
}
