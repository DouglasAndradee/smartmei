package repository

import (
	"context"
	"os"

	"github.com/DouglasAndradee/smartmei/database"
	"github.com/DouglasAndradee/smartmei/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
)

var session = database.Session()

// CountUser - Get numbers of users in the collection
func (r *Repository) CountUser(ctx context.Context) (int64, error) {
	r.mongo = session
	filter := bson.D{}
	count, err := r.mongo.Database(os.Getenv("DATABASE_NAME")).Collection("users").CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// CreateUser - Save the user in database
func (r *Repository) CreateUser(ctx context.Context, user entity.User) error {
	r.mongo = session
	_, err := r.mongo.Database(os.Getenv("DATABASE_NAME")).Collection("users").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

// CreateBook - Save the book in database
func (r *Repository) CreateBook(ctx context.Context, book entity.Book) error {
	r.mongo = session
	_, err := r.mongo.Database(os.Getenv("DATABASE_NAME")).Collection("users").InsertOne(ctx, book)
	if err != nil {
		return err
	}
	return nil
}

// LendBook - Save the loan in database
func (r *Repository) LendBook(ctx context.Context, loan entity.Loan) error {
	return nil
}

// ReturnBook - Returns the borrowed book to the owner
func (r *Repository) ReturnBook(ctx context.Context, book entity.Book) error {
	return nil
}
