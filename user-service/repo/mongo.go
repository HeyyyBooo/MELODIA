package repo

import (
    "context"
    "log"
    "os"
    "time"
	"fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func ConnectMongo() {
	fmt.Println("Connecting to MongoDB...")
    uri := os.Getenv("MONGO_URI")
    dbName := os.Getenv("MONGO_DB")
	fmt.Println("Using MONGO_URI:", uri)
	fmt.Println("Using MONGO_DB:", dbName)
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
	fmt.Println("Attempting MongoDB connection...")
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	fmt.Println("MongoDB connection attempt finished")
    if err != nil {
        log.Fatal("Mongo Connect Error: ", err)
    }
	fmt.Println("Pinging MongoDB...")
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Mongo Ping Error: ", err)
    }
	fmt.Println("MongoDB ping successful")
    Client = client
    DB = client.Database(dbName)
	fmt.Println("MongoDB connection established")
    log.Println("✅ Connected to MongoDB Atlas")
}
