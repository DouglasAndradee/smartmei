package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var once sync.Once
var instance *mongo.Client

//Session - This function connect in database and return the mongodb session if sucess.
func Session() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return mongogDB(os.Getenv("MONGODB_USER"), os.Getenv("MONGODB_PASSWORD"), os.Getenv("MONGODB_URI"))
}

func mongogDB(key, secret, uri string) *mongo.Client {
	once.Do(func() {
		settings := options.Client()
		settings.SetAuth(options.Credential{Username: key, Password: secret})
		settings.SetMaxPoolSize(200)
		settings.SetRetryReads(true)
		settings.SetRetryWrites(true)
		settings.SetReadPreference(readpref.Secondary())

		client, err := mongo.NewClient(settings.ApplyURI(uri))
		if err != nil {
			fmt.Printf("[ERROR] [MONGODB] - %s", err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			fmt.Printf("[ERROR] [MONGODB] - %s", err.Error())
		}

		defer cancel()

		err = client.Ping(ctx, readpref.Secondary())
		if err != nil {
			fmt.Printf("[ERROR] [MONGODB] - %s", err.Error())
		}
		instance = client
	})
	return instance
}
