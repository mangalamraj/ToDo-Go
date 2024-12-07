package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// func ConnectToMongo() {
// 	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
// 	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
// 	opts := options.Client().ApplyURI("mongodb+srv://mango26june:mango123@cluster0.ga9pq.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetServerAPIOptions(serverAPI)

// 	// Create a new client and connect to the server
// 	client, err := mongo.Connect(context.TODO(), opts)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer func() {
// 		if err = client.Disconnect(context.TODO()); err != nil {
// 			panic(err)
// 		}
// 	}()

//		// Send a ping to confirm a successful connection
//		if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
//			panic(err)
//		}
//		fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
//	}
func ConnectToMongo(uri string) error {
	// Set a timeout for connecting
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// Ping MongoDB to confirm the connection
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	// Set the global MongoClient
	MongoClient = client
	log.Println("Connected to MongoDB!")
	return nil
}

func GetCollection(database, collection string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatalf("MongoClient is not initialized. Call ConnectToMongo first.")
	}
	return MongoClient.Database(database).Collection(collection)
}
