package main

import (
    "log"
    "net"
    "os"

    "github.com/joho/godotenv"
    "melodia/user-service/handler"
    "melodia/user-service/repo"
    userpb "melodia/user-service/melodia/user-service/proto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

func main() {
    // Load env
    godotenv.Load(".env")
    port := os.Getenv("PORT")

    // Connect MongoDB
    repo.ConnectMongo()

    // Start gRPC server
    lis, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatal(err)
    }

    grpcServer := grpc.NewServer()
    userpb.RegisterUserServiceServer(grpcServer, &handler.UserServiceServer{})

    // Enable reflection for grpcurl testing
    reflection.Register(grpcServer)

    log.Printf("User Service running on port %s\n", port)
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatal(err)
    }
}
