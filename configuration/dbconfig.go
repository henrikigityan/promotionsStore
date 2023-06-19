package dbconfig

import (
	"context"
    "fmt"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


//Client instance
var DB *mongo.Client = connectDB()

func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
		log.Fatal(err)
    }
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Connected")
	return client
}

//getting database collections
func GetCollection(client *mongo.Client, collectionName string, databaseName string) *mongo.Collection {
    collection := client.Database(databaseName).Collection(collectionName)
    return collection
}
