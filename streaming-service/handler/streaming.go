package handler

import (
    "context"

    "melodia/streaming-service/proto"
    "melodia/streaming-service/repo"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type StreamingServer struct {
    proto.UnimplementedStreamingServiceServer
}

func (s *StreamingServer) UploadSong(ctx context.Context, req *proto.SongUploadRequest) (*proto.SongUploadResponse, error) {
    song := bson.M{
        "title":  req.Title,
        "artist": req.Artist,
        "album":  req.Album,
        "url":    req.Url,
    }

    _, err := repo.SongCollection.InsertOne(ctx, song)
    if err != nil {
        return nil, err
    }

    return &proto.SongUploadResponse{Message: "Song uploaded successfully!"}, nil
}

func (s *StreamingServer) GetSongList(ctx context.Context, req *proto.Empty) (*proto.SongListResponse, error) {
    cursor, err := repo.SongCollection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var songs []*proto.Song
    for cursor.Next(ctx) {
        var song bson.M
        cursor.Decode(&song)

        songs = append(songs, &proto.Song{
            Id:     song["_id"].(primitive.ObjectID).Hex(),
            Title:  song["title"].(string),
            Artist: song["artist"].(string),
            Album:  song["album"].(string),
            Url:    song["url"].(string),
        })
    }

    return &proto.SongListResponse{Songs: songs}, nil
}

func (s *StreamingServer) GetSongByID(ctx context.Context, req *proto.SongIDRequest) (*proto.SongResponse, error) {
    id, _ := primitive.ObjectIDFromHex(req.Id)
    var song bson.M
    err := repo.SongCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&song)
    if err != nil {
        return nil, err
    }

    result := &proto.Song{
        Id:     song["_id"].(primitive.ObjectID).Hex(),
        Title:  song["title"].(string),
        Artist: song["artist"].(string),
        Album:  song["album"].(string),
        Url:    song["url"].(string),
    }

    return &proto.SongResponse{Song: result}, nil
}
