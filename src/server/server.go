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
		messages := make([]kafka.Message, 10)
		for i := 0; i < 10; i++ {

			myproduct := models.GenerateProduct()

			res, err := json.Marshal(myproduct)
			if err != nil {
				log.Println("failed to marshal object ", err)
			}
			fmt.Printf("product: %s price: %d \n", myproduct.Name, myproduct.Price)
			fmt.Printf("bytes: %s", res)

			messages[i] = kafka.Message{Value: res}
		}

		for k, v := range messages {
			log.Printf("[%d]: %v", k, v.Value)
		}

		_, err = conn.WriteMessages(
			messages...,
		)

		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		time.Sleep(3 * time.Second)
	}

}
