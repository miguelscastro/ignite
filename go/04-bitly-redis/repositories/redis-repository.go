package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type repository struct {
	redisDB *redis.Client
}

type Repository interface {
	GetFullURL(ctx context.Context, code string) (string, error)
	SaveShortenedURL(ctx context.Context, _url string) (string, error)
}

func NewRepository(redisDB *redis.Client) Repository {
	return repository{redisDB}
}

func (repo repository) SaveShortenedURL(ctx context.Context, _url string) (string, error) {
	var code string
	for range 5 {
		code = genCode()
		if err := repo.redisDB.HGet(ctx, "shortener", code).Err(); err != nil {
			if errors.Is(err, redis.Nil) {
				break
			}
			return "", fmt.Errorf("failed to get code from shortener hashmap: %w", err)
		}

	}

	if err := repo.redisDB.HSet(ctx, "shortener", code, _url).Err(); err != nil {
		return "", fmt.Errorf("failed to set code in shortener hashmap: %w", err)
	}
	return code, nil
}

func (repo repository) GetFullURL(ctx context.Context, code string) (string, error) {
	fullURL, err := repo.redisDB.HGet(ctx, "shortener", code).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get code from shortener hashmap: %w", err)
	}
	return fullURL, nil
}
