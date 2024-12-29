package redispackage

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"

	"github.com/hentan/final_project/internal/config"
	"github.com/hentan/final_project/internal/logger"
	"github.com/hentan/final_project/internal/models"
	jsi "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	SetToCache(ctx context.Context, bookID int, book *models.Book, ttl time.Duration) error
	GetFromCache(ctx context.Context, bookID int) (*models.Book, error)
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisClient(cfg config.Config) *RedisCache {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
	})
	return &RedisCache{client: redisClient}
}

func (r *RedisCache) SetToCache(ctx context.Context, bookID int, book *models.Book, ttl time.Duration) error {
	log := logger.GetLogger()
	key := fmt.Sprintf("book:%d", bookID)
	data, err := jsi.Marshal(book)
	if err != nil {
		log.Error("не получилось распарсить book при парсинге redis: ", slog.String("err", err.Error()))
		return err
	}
	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *RedisCache) GetFromCache(ctx context.Context, bookID int) (*models.Book, error) {
	log := logger.GetLogger()
	key := fmt.Sprintf("book:%d", bookID)
	cachedBook, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			log.Info("Ключ не найден в кэше Redis", slog.String("key", key))
			return nil, nil
		}
		return nil, err
	}
	var book models.Book
	err = jsi.Unmarshal([]byte(cachedBook), &book)
	if err != nil {
		log.Error("не получается распарсить данные при получении кэша Redis", slog.String("err", err.Error()))
		return nil, err
	}
	return &book, nil

}

func (r *RedisCache) getCacheHitRatio(ctx context.Context) (float64, error) {
	log := logger.GetLogger()
	info, err := r.client.Info(ctx, "stats").Result()
	if err != nil {
		log.Error("не получается вызвать метод Info для Redis", slog.String("err", err.Error()))
		return 0, err
	}
	var hits, misses int64
	for _, line := range strings.Split(info, "\n") {
		if strings.HasPrefix(line, "keyspace_hits") {
			hits, _ = strconv.ParseInt(strings.TrimPrefix(line, "keyspace_hits:"), 10, 64)
		}
		if strings.HasPrefix(line, "keyspace_misses") {
			misses, _ = strconv.ParseInt(strings.TrimPrefix(line, "keyspace_misses:"), 10, 64)
		}
	}
	if hits+misses == 0 {
		return 0, nil
	}

	return (float64(hits) / float64(hits+misses)) * 100, nil
}
