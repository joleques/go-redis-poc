package useCase

import (
	"fmt"
	"github.com/joleques/go-redis-poc/src/application"
	"github.com/joleques/go-redis-poc/src/infra/redis"
	"strconv"
)

func CreateOrder(identifier string, products []string) error {
	order, err := application.NewOrder(identifier, products)
	if err != nil {
		return err
	}

	redisClient, errRedis := redis.NewRedis()
	if errRedis != nil {
		return errRedis
	}

	for i, product := range order.Products {
		id := order.Identifier + "_" + strconv.Itoa(i)
		fmt.Println(i, id, product)
		redisClient.Set(id, product, 0)
	}

	return nil
}
