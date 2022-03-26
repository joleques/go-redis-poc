package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"os"
	"sync"
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

func (client clientRedisAdapter) Set(key string, value string, expiration time.Duration, wg *sync.WaitGroup) {

	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	logger = logger.With(zap.String("app", "go-redis")).With(zap.String("environment", "test")).With(zap.String("key", key)).With(zap.String("value", value))

	err := client.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		panic(err)
	}
	logger.Info("Insert successfully")
	wg.Done()
}
