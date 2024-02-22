package adapter

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisAdapter interface {
	SendCache(key string, value string) error
	GetCache(key string) (string, error)
}

type redisAdapterLayer struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedisAdapter(rdb *redis.Client) RedisAdapter {
	return &redisAdapterLayer{rdb: rdb, ctx: context.Background()}
}

func (layer redisAdapterLayer) SendCache(key string, value string) error {
	err := layer.rdb.Set(layer.ctx, key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
		fmt.Println("error al enviar cache", err)
		return err
	}
	return nil
}

func (layer redisAdapterLayer) GetCache(key string) (string, error) {
	value, err := layer.rdb.Get(layer.ctx, key).Result()
	if err != nil {
		fmt.Println("Llave no encontrada")
		return "", err
	}
	return value, nil
}
