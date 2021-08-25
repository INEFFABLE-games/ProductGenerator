package server

import (
	"ProductsGenerator/src/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"time"
)

func StartServer() {
	ip := os.Getenv("ZADDRES")

	topic := "products"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", ip, topic, partition)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	for {

		myproduct := models.GenerateProduct()

		res, err := json.Marshal(myproduct)
		if err != nil {
			log.Println("failed to marshal object ", err)
		}
		fmt.Printf("product: %s price: %d \n", myproduct.Name, myproduct.Price)
		fmt.Printf("bytes: %s", res)

		_, err = conn.WriteMessages(
			kafka.Message{Value: res},
		)

		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		fmt.Println(res)
		time.Sleep(3 * time.Second)
	}

}
