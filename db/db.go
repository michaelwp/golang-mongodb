package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func MongoDB() *mongo.Database{
	user := loadDBCred("DB_USER")
	pass := loadDBCred("DB_PASS")
	db := loadDBCred("DB_NAME")
	host := loadDBCred("DB_HOST")
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

func loadDBCred(key string) string {
	err := godotenv.Load(".env")
	if err != nil {log.Fatal(err)}

	return os.Getenv(key)
}
