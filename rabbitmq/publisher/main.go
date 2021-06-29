package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://myuser:password@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"events", // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		fmt.Println(err)
	}

	for i := 1; i < 30; i++ {
		time.Sleep(5 * time.Second)

		message := fmt.Sprintf("Message n. %d: %s", i, randomString(6))
		// attempt to publish a message to the queue!
		err = ch.Publish(
			"events",
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Successfully Published Message to Queue: ", message)
	}
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
