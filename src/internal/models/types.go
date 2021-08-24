package models

import (
	"math/rand"
	"strconv"
)

type product struct {
	Name string `json:"name" bson:"name"`
	Price uint64 `json:"price" bson:"price"`
}

func GenerateProduct() product{
	return product{Name: strconv.Itoa(rand.Int()),Price: rand.Uint64()}
}