package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongo() {
	fmt.Println("mongo connection start")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost/?replicaSet=yourReplicaSet")
	var err error
	MongoClient, err = mongo.Connect(ctx, clientOptions)
	fmt.Println(err)

	if err != nil {
		fmt.Println("fail mongo connection start")
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
}
