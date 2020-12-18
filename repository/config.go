package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository -
type Repository struct {
	Session *mongo.Client
}

var databaseName = "smartmei"
