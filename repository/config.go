package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository - It's a instance of mongo client
type Repository struct {
	Session *mongo.Client
}

var databaseName = "loanbooks"
