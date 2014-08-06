package config

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

var RedisConnection redis.Conn

func SetupRedis() {
	RedisConnection, err := redis.Dial("tcp", ":6379")

	if err != nil {
		log.Fatal(err)
	}

	RedisConnection.Do("INFO")
}

func ShutdownRedis() {
	RedisConnection.Close()
}
