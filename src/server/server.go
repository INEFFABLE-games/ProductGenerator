package server

import (
	"ProductsGenerator/src/internal/models"
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"time"
)

//kafka settings
const (
	topic = "products"
	partition = 0
)

func StartServer() error {

	ip := os.Getenv("ZADDRES")

	conn, err := kafka.DialLeader(context.Background(), "tcp", ip + ":9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	for {
		myproduct := models.GenerateProduct()
		//fmt.Printf("product: %s price: %d \n", myproduct.Name, myproduct.Price)

		var b bytes.Buffer
		enc := gob.NewEncoder(&b)

		err := enc.Encode(myproduct)
		if err != nil {
			log.Fatal("encode error:", err)
		}

		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte(b.Bytes())},
		)

		fmt.Printf("message was send: %s %d",myproduct.Name,myproduct.Price )

		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		if err := conn.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}

		time.Sleep(3 * time.Second)
	}

	return nil
}
