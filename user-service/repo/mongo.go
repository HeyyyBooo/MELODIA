package repo

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func ConnectMongo() {
    uri := os.Getenv("MONGO_URI")
    dbName := os.Getenv("MONGO_DB")

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("Mongo Connect Error: ", err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("Mongo Ping Error: ", err)
    }

    Client = client
    DB = client.Database(dbName)
    log.Println("✅ Connected to MongoDB Atlas")
}
