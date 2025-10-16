package handler

import (
    "context"
    "fmt"
    "log"

    "melodia/user-service/models"
    "melodia/user-service/repo"
    userpb "melodia/user-service/melodia/user-service/proto"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "golang.org/x/crypto/bcrypt"
)

type UserServiceServer struct {
    userpb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
    collection := repo.DB.Collection("users")

    // Check if email exists
    var existing models.User
    err := collection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&existing)
    if err == nil {
        return &userpb.RegisterResponse{
            Success: false,
            Message: "Email already registered",
        }, nil
    }

    // Hash password
    hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

    user := models.User{
        ID:       primitive.NewObjectID(),
        Username: req.Username,
        Email:    req.Email,
        Password: string(hashed),
    }

    _, err = collection.InsertOne(ctx, user)
    if err != nil {
        return &userpb.RegisterResponse{
            Success: false,
            Message: "Error creating user",
        }, err
    }

    log.Println("New user registered:", req.Email)
    return &userpb.RegisterResponse{
        Success: true,
        Message: "User registered successfully",
        UserId:  user.ID.Hex(),
    }, nil
}

func (s *UserServiceServer) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
    collection := repo.DB.Collection("users")

    var user models.User
    err := collection.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
    if err != nil {
        return &userpb.LoginResponse{
            Success: false,
            Token:   "",
        }, nil
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        return &userpb.LoginResponse{
            Success: false,
            Token:   "",
        }, nil
    }

    // For now, dummy token
    token := fmt.Sprintf("dummy-token-for-%s", user.ID.Hex())
    return &userpb.LoginResponse{
        Success: true,
        Token:   token,
    }, nil
}
