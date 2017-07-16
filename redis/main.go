package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

var (
	c redis.Conn
)

func getRecord(key string) string {

	item, err := redis.String(c.Do("GET", key))

	if err == redis.ErrNil {
		log.Fatal("Record not found in redis")
		return ""
	} else if err != nil {
		PanicOnError(err)
	}

	return item
}

func main() {

	var err error

	c, err = redis.Dial("tcp", ":6379")
	PanicOnError(err)

	c.Do("SET", "message1", "Hello World")

	item := getRecord("message1")
	fmt.Println(item)

	item = getRecord("name")
	fmt.Println(item)

	// Set redis expires in 5 seconds
	/*result, err := redis.String(c.Do("SET", "driver", 5715, "EX", 5))
	PanicOnError(err)

	if result != "OK" {
		panic("not set data to redis" + result)
	}

	fmt.Println(getRecord("driver"))
	*/

	// MGET

	keys := []interface{}{"name", "message1"}

	results, err := redis.Strings(c.Do("MGET", keys...))
	PanicOnError(err)

	fmt.Println(results)

	defer c.Close()
}

// PanicOnError handles errors
func PanicOnError(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
