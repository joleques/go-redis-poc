package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type clientRedisAdapter struct {
	Addr   string
	DB     int
	Client *redis.Client
}

var ctx = context.Background()

func NewRedis() (*clientRedisAdapter, error) {
	clientRedis := clientRedisAdapter{Addr: "localhost:6379", DB: 0}
	clientRedis.connect()
	_, err := clientRedis.Client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &clientRedis, nil
}

func (client *clientRedisAdapter) connect() {
	rdb := redis.NewClient(&redis.Options{
		Addr: client.Addr,
		DB:   client.DB, // use default DB
	})

	client.Client = rdb
}

func (client clientRedisAdapter) Set(key string, value string, expiration time.Duration) {
	err := client.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert successfully")
}
