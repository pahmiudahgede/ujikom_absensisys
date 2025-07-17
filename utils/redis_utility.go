package utils

import (
	"absensibe/config"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisUtil struct{}

func NewRedisUtil() *RedisUtil {
	return &RedisUtil{}
}

func (r *RedisUtil) Set(key string, value interface{}, ttl ...time.Duration) error {
	var expiration time.Duration
	if len(ttl) > 0 {
		expiration = ttl[0]
	}

	var val interface{}
	switch v := value.(type) {
	case string, int, int64, float64, bool:
		val = v
	default:

		jsonData, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("failed to marshal value: %v", err)
		}
		val = string(jsonData)
	}

	return config.Redis.Set(ctx, key, val, expiration).Err()
}

func (r *RedisUtil) Get(key string) (string, error) {
	val, err := config.Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key '%s' not found", key)
	}
	return val, err
}

func (r *RedisUtil) GetJSON(key string, dest interface{}) error {
	val, err := r.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

func (r *RedisUtil) Delete(keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	return config.Redis.Del(ctx, keys...).Err()
}

func (r *RedisUtil) Exists(key string) bool {
	count := config.Redis.Exists(ctx, key).Val()
	return count > 0
}

func (r *RedisUtil) TTL(key string) (time.Duration, error) {
	return config.Redis.TTL(ctx, key).Result()
}

func (r *RedisUtil) Expire(key string, ttl time.Duration) error {
	return config.Redis.Expire(ctx, key, ttl).Err()
}

func (r *RedisUtil) HSet(key, field string, value interface{}) error {
	return config.Redis.HSet(ctx, key, field, value).Err()
}

func (r *RedisUtil) HGet(key, field string) (string, error) {
	val, err := config.Redis.HGet(ctx, key, field).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("field '%s' not found in hash '%s'", field, key)
	}
	return val, err
}

func (r *RedisUtil) HGetAll(key string) (map[string]string, error) {
	return config.Redis.HGetAll(ctx, key).Result()
}

func (r *RedisUtil) HDel(key string, fields ...string) error {
	return config.Redis.HDel(ctx, key, fields...).Err()
}

func (r *RedisUtil) HExists(key, field string) bool {
	return config.Redis.HExists(ctx, key, field).Val()
}

func (r *RedisUtil) LPush(key string, values ...interface{}) error {
	return config.Redis.LPush(ctx, key, values...).Err()
}

func (r *RedisUtil) RPush(key string, values ...interface{}) error {
	return config.Redis.RPush(ctx, key, values...).Err()
}

func (r *RedisUtil) LPop(key string) (string, error) {
	val, err := config.Redis.LPop(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("list '%s' is empty", key)
	}
	return val, err
}

func (r *RedisUtil) RPop(key string) (string, error) {
	val, err := config.Redis.RPop(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("list '%s' is empty", key)
	}
	return val, err
}

func (r *RedisUtil) LLen(key string) int64 {
	return config.Redis.LLen(ctx, key).Val()
}

func (r *RedisUtil) LRange(key string, start, stop int64) ([]string, error) {
	return config.Redis.LRange(ctx, key, start, stop).Result()
}

func (r *RedisUtil) SAdd(key string, members ...interface{}) error {
	return config.Redis.SAdd(ctx, key, members...).Err()
}

func (r *RedisUtil) SMembers(key string) ([]string, error) {
	return config.Redis.SMembers(ctx, key).Result()
}

func (r *RedisUtil) SIsMember(key string, member interface{}) bool {
	return config.Redis.SIsMember(ctx, key, member).Val()
}

func (r *RedisUtil) SRem(key string, members ...interface{}) error {
	return config.Redis.SRem(ctx, key, members...).Err()
}

func (r *RedisUtil) SCard(key string) int64 {
	return config.Redis.SCard(ctx, key).Val()
}

func (r *RedisUtil) Keys(pattern string) ([]string, error) {
	return config.Redis.Keys(ctx, pattern).Result()
}

func (r *RedisUtil) FlushPattern(pattern string) error {
	keys, err := r.Keys(pattern)
	if err != nil {
		return err
	}
	if len(keys) > 0 {
		return r.Delete(keys...)
	}
	return nil
}

func (r *RedisUtil) Ping() error {
	return config.Redis.Ping(ctx).Err()
}

func (r *RedisUtil) FlushDB() error {
	return config.Redis.FlushDB(ctx).Err()
}

func (r *RedisUtil) SetCache(key string, data interface{}, ttl time.Duration) error {
	return r.Set(key, data, ttl)
}

func (r *RedisUtil) GetCache(key string, dest interface{}) error {
	return r.GetJSON(key, dest)
}

func (r *RedisUtil) GetOrSetCache(key string, dest interface{}, ttl time.Duration, fetchFunc func() (interface{}, error)) error {

	err := r.GetCache(key, dest)
	if err == nil {
		return nil
	}

	data, err := fetchFunc()
	if err != nil {
		return fmt.Errorf("failed to fetch data: %v", err)
	}

	if err := r.SetCache(key, data, ttl); err != nil {

		fmt.Printf("Warning: failed to cache data for key '%s': %v\n", key, err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal fetched data: %v", err)
	}

	return json.Unmarshal(jsonData, dest)
}

func (r *RedisUtil) IncrementCounter(key string, ttl time.Duration) (int64, error) {

	val, err := config.Redis.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if val == 1 {
		if err := r.Expire(key, ttl); err != nil {
			return val, fmt.Errorf("failed to set TTL: %v", err)
		}
	}

	return val, nil
}

var Redis = NewRedisUtil()

const (
	TTL_SHORT  = 5 * time.Minute
	TTL_MEDIUM = 1 * time.Hour
	TTL_LONG   = 24 * time.Hour
	TTL_WEEK   = 7 * 24 * time.Hour
)
