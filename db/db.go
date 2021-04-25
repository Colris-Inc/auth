package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	ConnectionString string
}

//connectDB connects with the DB and does a ping to test the connection
//if the connection is successful, it returns a mongo client for further operations
func connectDB(connectionString string) (*mongo.Client, error) {
	clientOpts := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err = client.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return client, err
}

//disconnectDB disconnects the current
func disconnectDB(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
}
