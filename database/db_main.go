package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client   *mongo.Client
	users    *mongo.Collection
	verify   *mongo.Collection
	question *mongo.Collection
}

func Connect() *DB {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:admin2boogaloo@cluster0.0mieo.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Print(err)
		log.Print("\nDB connection failed in database package")
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	return &DB{
		client:   client,
		users:    client.Database("cinderellaApp").Collection("users"),
		verify:   client.Database("cinderellaApp").Collection("verify"),
		question: client.Database("cinderellaApp").Collection("question"),
	}
}

// Disconnect to MongoDB - could use, on further inspection looks like mongo is good at closing itself
// func (db *DB) closeDB() {
// 	err := db.client.Disconnect(context.Background())
// 	if err != nil {
// 		log.Print(err)
//      log.Print("\nunable to disconnect from DB in database package\n")
//      return nil
// 	}
// }
