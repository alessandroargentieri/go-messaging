package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	log.Println("Publisher started")

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

	log.Println("Connected to Redis server")

	for i := 0; i < 30; i++ {
		// Sleep random time
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(4)
		time.Sleep(time.Duration(n) * time.Second)
		// publish message on stream
		err = publishTicketReceivedEvent(redisClient)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func publishTicketReceivedEvent(client *redis.Client) error {
	log.Println("Publishing event to Redis")

	err := client.XAdd(&redis.XAddArgs{
		Stream:       "tickets",
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "",
		Values: map[string]interface{}{
			"whatHappened": string("ticket received"),
			"ticketID":     int(rand.Intn(100000000)),
			"ticketData":   string("some ticket data"),
		},
	}).Err()

	return err
}
