package db

import (
	"context"
	"fmt"
	"gin-api/internal/config"
	"github.com/go-redis/redis/v8"
)

// ConnectRedis 连接到Redis并返回redis实例
func ConnectRedis(conf config.RedisConfig) *redis.Client {
	// 创建一个上下文
	ctx := context.Background()

	// 创建Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Errorf("failed to connect to Redis: %w", err))
	}
	fmt.Println("Redis连接成功")
	return rdb
}
