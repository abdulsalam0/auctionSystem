package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// create a client
func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

func get(client *redis.Client) error {
	value, err := client.Get("server1").Result()
	if err != nil {
		return err
	}
	fmt.Println("server 1 port", value)

	portvalue, err := client.Get("server2").Result()
	if err != nil {
		return err
	}
	fmt.Println("server 2 port", portvalue)
	return nil
}

func viewList(client *redis.Client) error {
	value, err := client.LRange("server_list", 0, -1).Result()
	if err != nil {
		return err
	}
	fmt.Println(value)
	return nil
}

func main() {
	fmt.Println("we are up and running")

	client := createClient()

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	for {
		err := get(client)
		if err != nil {
			fmt.Println(err)
		}
		err1 := viewList(client)
		if err != nil {
			fmt.Println(err1)
		}
		time.Sleep(time.Millisecond * 5000)
	}
}
