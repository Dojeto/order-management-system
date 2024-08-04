package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	URL string
}

type Store struct {
	Config
	client *mongo.Client
}

func NewStore(config Config) *Store {
	return &Store{
		Config: config,
	}
}

func (s Store) Create(ctx context.Context) error {
	var err error

	clientOptions := options.Client().ApplyURI(s.Config.URL)

	s.client, err = mongo.Connect(ctx, clientOptions)

	if err != nil {
		return err
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = s.client.Ping(pingCtx, readpref.Primary())

	if err != nil {
		return err
	}
	fmt.Println("Connected to MongoDB!")

	return nil
}
