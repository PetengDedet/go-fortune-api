package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetMongoClient(connString string) *mongo.Client {
	clientOption := options.Client().ApplyURI(connString)

	var err error
	client, err = mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetDB(client *mongo.Client, dbname string) *mongo.Database {
	return client.Database(dbname)
}

func CloseMongoConnection() {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// TODO optional you can log your closed MongoDB client
	fmt.Println("Connection to MongoDB closed.")
}
