package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func mongoDB() *mongo.Database{
	user := "michaelwp"
	pass := "michaelwp"
	db := "db_user"
	host := "cluster0-jjufo.mongodb.net"
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		user, pass, host, db)
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil { log.Fatal(err) }

	err = client.Ping(ctx, nil)
	if err != nil { log.Fatal(err) }

	fmt.Println("Connected to MongoDB!")

	return client.Database(db)
}
