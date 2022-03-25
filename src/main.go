package main

import (
	"context"
	"fmt"
	"github.com/joleques/go-redis-poc/src/useCase"
)

var ctx = context.Background()

func main() {

	err := useCase.CreateOrder("789", []string{"mesa", "cadeira"})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Process finish....")
}
