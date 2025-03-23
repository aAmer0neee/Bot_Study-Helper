package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	RedisClient *redis.Client
}

func NewRedisClient(port string) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     port, // Адрес Redis-сервера (по умолчанию 6379)
		Password: "",   // Пароль Redis (если нет, оставляем пустым)
		DB:       0,    // Используемая база данных Redis
	})

	// Проверка подключения
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Не удалось подключиться к Redis: %v", err)
	}
	fmt.Println("Подключено к Redis!")
	return &Cache{RedisClient: rdb}
}

func (r *Cache) AddRecord( key string, value string, exp time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := r.RedisClient.Set(ctx, key, value, exp).Err(); err != nil {
		log.Fatalf(err.Error())
	}
	res , _ := r.GetRecord(key)
	log.Println("\n\n\n",res,"\n\n\n")
}
func (r *Cache) GetRecord( key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	value := r.RedisClient.Get(ctx, key)
	res, err := value.Result()
	if err == redis.Nil {
		return "", fmt.Errorf("no value")
	}
	return res, nil
}
