package repo

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var SongCollection *mongo.Collection

func InitMongo() {
    uri := os.Getenv("MONGO_URI")
    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal("MongoDB Connection Error: ", err)
    }

    db := os.Getenv("MONGO_DB")
    fmt.Println("✅ Connected to MongoDB:", db)
    SongCollection = client.Database(db).Collection("songs")
}
