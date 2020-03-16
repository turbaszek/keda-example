package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

// Create Redis client
func makeClient() *redis.Client {
	options := &redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "",
		DB:       0,
	}

	return redis.NewClient(options)
}

// Publish messages to Redis list
func Publish() (int64, error) {
	client := makeClient()

	// Create sample messages
	values := make([]string, 100)
	for i := range values {
		values[i] = fmt.Sprintf("Message #%d\n", i)
	}

	cmd := client.RPush("default", values)
	if cmd.Err() != nil {
		return 0, cmd.Err()
	}
	return cmd.Result()
}

// Drain trim Redis list to 0
func Drain() (string, error) {
	client := makeClient()
	cmd := client.LTrim("default", 1, 0)

	if cmd.Err() != nil {
		return "error", cmd.Err()
	}
	return cmd.Result()
}

// GetListLength returns length of redis list
func GetListLength() (int64, error) {
	client := makeClient()
	cmd := client.LLen("default")

	if cmd.Err() != nil {
		return 0, cmd.Err()
	}
	return cmd.Result()
}
