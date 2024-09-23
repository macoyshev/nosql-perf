package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/macoyshev/nosql-perf/databases/redis"
	"github.com/macoyshev/nosql-perf/timer"
)

func main() {
	redisClient := redis.NewClient()
	defer redisClient.Close()

	file, err := os.ReadFile("./cases/rw-json/data.json")
	if err != nil {
		panic(err)
	}

	endTimer := timer.PrintTime("redis write json")
	status := redisClient.Set(context.TODO(), "json", file, time.Duration(1000000))
	endTimer()

	res, err := status.Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("redis write json res:", res)
}
