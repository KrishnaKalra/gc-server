package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	playerDb = "playerdb"
	player   = "players"
)

var (
	PlayerCollection *mongo.Collection
	Gcsb             *mongo.Database
)

func InitCore() {
	// err := godotenv.Load()

	// if err != nil {
	// 	fmt.Println("Error loading .env")
	// 	return
	// }

	DB_URI := os.Getenv("DB")

	clientOptions := options.Client().ApplyURI(DB_URI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	Gcsb = client.Database(playerDb)
	PlayerCollection = Gcsb.Collection(player)

	fmt.Println("PlayerDB/Player Collection is Ready. -1")
}
