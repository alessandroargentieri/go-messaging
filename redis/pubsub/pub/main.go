package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

// User is a struct representing newly registered users
type User struct {
	Username string
	Email    string
}

// MarshalBinary encodes the struct into a binary blob
// Here I cheat and use regular json :)
func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalBinary decodes the struct into a User
func (u User) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &u); err != nil {
		return err
	}
	return nil
}

// names Some Non-Random name lists used to generate Random Users
var names []string = []string{"Jasper", "Johan", "Edward", "Niel", "Percy", "Adam", "Grape", "Sam", "Redis", "Jennifer", "Jessica", "Angelica", "Amber", "Watch"}

// lastnames Some Non-Random name lists used to generate Random Users
var lastnames []string = []string{"Ericsson", "Redisson", "Edisson", "Tesla", "Bolmer", "Andersson", "Sword", "Fish", "Coder"}

// emailProviders Some Non-Random email lists used to generate Random Users
var emailProviders []string = []string{"Hotmail.com", "Gmail.com", "Awesomeness.com", "Redis.com"}

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

	fmt.Println("Publisher inited on topic 'new_users'")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	// Loop and randomly generate users on a random timer
	for {
		// Publish a generated user to the new_users channel
		err := redisClient.Publish("new_users", GenerateRandomUser()).Err()
		if err != nil {
			panic(err)
		}
		fmt.Println("a new user has been published")
		// Sleep random time
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(4)
		time.Sleep(time.Duration(n) * time.Second)
	}

}

// GenerateRandomUser creates a random user, dont care too much about this.
func GenerateRandomUser() User {
	rand.Seed(time.Now().UnixNano())
	nameMax := len(names)
	lastnameMax := len(lastnames)
	emailProviderMax := len(emailProviders)

	nameIndex := rand.Intn(nameMax-1) + 1
	lastnameIndex := rand.Intn(lastnameMax-1) + 1
	emailIndex := rand.Intn(emailProviderMax-1) + 1

	return User{
		Username: names[nameIndex] + " " + lastnames[lastnameIndex],
		Email:    strings.ToLower(names[nameIndex] + "." + lastnames[lastnameIndex] + "@" + emailProviders[emailIndex]),
	}
}
