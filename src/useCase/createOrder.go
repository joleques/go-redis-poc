package useCase

import (
	"github.com/joleques/go-redis-poc/src/application"
	"github.com/joleques/go-redis-poc/src/infra/redis"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func CreateOrder(identifier string, products []string) error {
	order, err := application.NewOrder(identifier, products)
	if err != nil {
		return err
	}

	redisClient, errRedis := redis.NewRedis()
	if errRedis != nil {
		return errRedis
	}

	wg.Add(order.TotalProducts())

	for i, product := range order.Products {
		id := order.Identifier + "_" + strconv.Itoa(i)
		go redisClient.Set(id, product, 0, &wg)
	}

	wg.Wait()
	return nil
}
