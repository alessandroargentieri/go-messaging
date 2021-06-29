package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	// Create a new Redis Client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // We connect to host redis, thats what the hostname of the redis service is set to in the docker-compose
		Password: "superSecret",    // The password IF set in the redis Config file
		DB:       0,
	})
	// Ping the Redis server and check if any errors occured
	err := redisClient.Ping().Err()
	if err != nil {
		// Sleep for 3 seconds and wait for Redis to initialize
		time.Sleep(3 * time.Second)
		err := redisClient.Ping().Err()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Subscriber inited on topic 'new_users'")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// Subscribe to the Topic given
	topic := redisClient.Subscribe("new_users")
	// Get the Channel to use
	channel := topic.Channel()
	// Iterate any messages sent on the channel
	for msg := range channel {
		fmt.Println("Read a new user:", msg.Payload)
	}
}
