package main

import (
	"context"
	"log"

	"github.com/yervsil/user_service/config"
	"github.com/yervsil/user_service/internal/server"
	"github.com/yervsil/user_service/pkg/logger"
	"github.com/yervsil/user_service/pkg/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	log.Println("Starting user microservice")

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.New(cfg.Env)

	mongoDBConn, err := mongodb.NewClient(cfg.Mongo.URI, cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		appLogger.Fatal("cannot connect mongodb", err)
	}
	
	db := mongoDBConn.Database(cfg.Mongo.Name)
	appLogger.Info("MongoDB connected: %v", mongoDBConn.NumberSessionsInProgress())

	collection := db.Collection("user")
	_, err = collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
        {Keys: bson.M{"username": 1}, Options: options.Index().SetUnique(true)},
        {Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true)},
    })
    if err != nil {
		appLogger.Fatal("cannot create indexes to the fields", err)
    }

	redisClient := redis.NewRedisClient(cfg)
	defer redisClient.Close()
	appLogger.Info("Redis connected")

	s := server.NewServer(appLogger, cfg, db, redisClient)
	appLogger.Fatal(s.Run())
}