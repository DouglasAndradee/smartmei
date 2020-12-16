package repository

import "go.mongodb.org/mongo-driver/mongo"

// Repository -
type Repository struct {
	mongo *mongo.Client
}
