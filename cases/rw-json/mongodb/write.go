package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/macoyshev/nosql-perf/databases/mongodb"
	"github.com/macoyshev/nosql-perf/timer"
)

func main() {
	mongoClient := mongodb.NewClient()
	defer mongoClient.Disconnect(context.TODO())

	file, err := os.ReadFile("./cases/rw-json/data.json")
	if err != nil {
		panic(err)
	}

	collection := mongoClient.Database("mongo").Collection("json")
	var data any
	err = json.Unmarshal(file, &data)
	if err != nil {
		panic(err)
	}

	endTimer := timer.PrintTime("mongodb json write")
	res, err := collection.InsertOne(context.TODO(), data)
	endTimer()

	if err != nil {
		panic(err)
	}
	fmt.Println("mongodb write json res:", res.InsertedID)
}
