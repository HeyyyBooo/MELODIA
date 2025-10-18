package main

import (
    "fmt"
    "log"
    "net"
    "os"

    "github.com/joho/godotenv"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"

    "melodia/streaming-service/handler"
    "melodia/streaming-service/proto"
    "melodia/streaming-service/repo"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    repo.InitMongo()

    port := os.Getenv("PORT")
    lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    proto.RegisterStreamingServiceServer(grpcServer, &handler.StreamingServer{})
    reflection.Register(grpcServer)

    log.Printf("🎵 Streaming Service running on port %s", port)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
