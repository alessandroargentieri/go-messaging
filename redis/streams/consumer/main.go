package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/rs/xid"
)

func main() {
	log.Println("Consumer started")

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

	subject := "tickets"
	consumersGroup := "tickets-consumer-group"

	err = redisClient.XGroupCreate(subject, consumersGroup, "0").Err()
	if err != nil {
		log.Println(err)
	}

	uniqueID := xid.New().String()

	for {
		entries, err := redisClient.XReadGroup(&redis.XReadGroupArgs{
			Group:    consumersGroup,
			Consumer: uniqueID,
			Streams:  []string{subject, ">"},
			Count:    2,
			Block:    0,
			NoAck:    false,
		}).Result()
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < len(entries[0].Messages); i++ {
			messageID := entries[0].Messages[i].ID
			values := entries[0].Messages[i].Values
			eventDescription := fmt.Sprintf("%v", values["whatHappened"])
			ticketID := fmt.Sprintf("%v", values["ticketID"])
			ticketData := fmt.Sprintf("%v", values["ticketData"])

			if eventDescription == "ticket received" {
				err := handleNewTicket(ticketID, ticketData)
				if err != nil {
					log.Fatal(err)
				}
				// sending acknowledgement to REDIS
				redisClient.XAck(subject, consumersGroup, messageID)
			}
		}
	}
}

func handleNewTicket(ticketID string, ticketData string) error {
	log.Printf("Handling new ticket id : %s data %s\n", ticketID, ticketData)
	return nil
}
