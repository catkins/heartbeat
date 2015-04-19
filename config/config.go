package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"gopkg.in/redis.v2"
)

// Configuration runtime config options for heartbeat service
type Configuration struct {
	RedisAddress      string
	RedisPassword     string
	RedisDatabase     int64
	HeartbeatMessage  string
	HeartbeatChannel  string
	HeartbeatInterval int64
}

func init() {
	godotenv.Load()
}

// Load load configuration from yaml
func Load() Configuration {
	redisURL := getEnvWithFallback("REDIS_URL", "localhost:6379")
	// strip out heroku style redis address prefixes
	redisURL = strings.Replace(redisURL, "redis://", "", 1)

	return Configuration{
		RedisAddress:      redisURL,
		RedisPassword:     getEnvWithFallback("REDIS_PASSWORD", ""),
		RedisDatabase:     getIntEnvWithFallback("REDIS_DATABASE", 0),
		HeartbeatChannel:  getEnvWithFallback("HEARTBEAT_CHANNEL", "heartbeat"),
		HeartbeatMessage:  getEnvWithFallback("HEARTBEAT_MESSAGE", ""),
		HeartbeatInterval: getIntEnvWithFallback("HEARTBEAT_INTERVAL", 1),
	}
}

// RedisOptions returns new go-redis connection options
func (c *Configuration) RedisOptions() redis.Options {
	return redis.Options{
		Addr:     c.RedisAddress,
		DB:       c.RedisDatabase,
		Password: c.RedisPassword,
	}
}

func getIntEnvWithFallback(key string, fallback int64) int64 {
	str := os.Getenv(key)
	if len(str) == 0 {
		return fallback
	}

	value, err := strconv.Atoi(str)

	if err != nil {
		panic("Invalid environment value for " + key + ". Expected integer.")
	}

	return int64(value)
}

func getEnvWithFallback(key string, fallback string) string {
	str := os.Getenv(key)

	if len(str) == 0 {
		return fallback
	}

	return str
}
