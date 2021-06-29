package main

import (
	"fmt"
	"time"

	"encoding/json"

	"github.com/go-redis/redis"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
	if err != nil {
		fmt.Println(err)
	}

	// SET ITEM
	err = client.Set("id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Author set in Redis:", string(json))
	}

	// GET ITEM
	val, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println("Error while fetching id id1234", err)
	}
	fmt.Println("Author fetched from Redis:", val)

	// DELETE ITEM
	client.Del("id1234")

	// TRY TO GET DELETED ITEM
	valAgain, err := client.Get("id1234").Result()
	if err != nil {
		fmt.Println("Error while fetching id id1234", err)
	}
	fmt.Println("Author fetched from Redis again:", valAgain)

	// SET ITEM SPECIFYING THE EXPIRATION TIME
	client.SetNX("exxppp", "expiringItem", 5*time.Millisecond)

	// SLEEP THE GOROUTINE SO THE ITEM EXPIRES
	time.Sleep(10 * time.Millisecond)

	// TRY TO GET EXPIRED ITEM
	valExp, err := client.Get("exxppp").Result()
	if err != nil {
		fmt.Println("Error while fetching id exxppp", err)
	}
	fmt.Println("Author fetched from Redis:", valExp)

}
