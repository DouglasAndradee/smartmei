package repository

import (
	"context"

	"github.com/DouglasAndradee/smartmei/domain"
	"go.mongodb.org/mongo-driver/bson"
)

// InsertUser -
func (r *Repository) InsertUser(ctx context.Context, user domain.User) (*domain.User, error) {

	if err := user.Valid(); err != nil {
		return nil, err
	}

	user.DefaultFields()

	result, err := r.Session.Database(databaseName).Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID
	return &user, nil
}

// CountUser -
func (r *Repository) CountUser(ctx context.Context) (*int64, error) {
	filter := bson.M{}
	count, err := r.Session.Database(databaseName).Collection("users").CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &count, nil
}

//GetUser -
func (r *Repository) GetUser(ctx context.Context, filter interface{}) (*domain.User, error) {
	result := domain.User{}
	err := r.Session.Database(databaseName).Collection("users").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// FoundBook -
func (r *Repository) FoundBook(ctx context.Context, filter interface{}) (interface{}, error) {
	result := domain.User{}
	err := r.Session.Database(databaseName).Collection("users").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// AddBook -
func (r *Repository) AddBook(ctx context.Context, filter interface{}, book domain.Book) (interface{}, error) {

	if err := book.Valid(); err != nil {
		return nil, err
	}

	update := bson.M{"$push": bson.M{"collection": book}}
	_, err := r.Session.Database(databaseName).Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// LendBook -
func (r *Repository) LendBook(ctx context.Context, filter interface{}, loan domain.Loan) (interface{}, error) {

	if err := loan.Valid(); err != nil {
		return nil, err
	}

	update := bson.M{"$addToSet": bson.M{"lent_books": loan, "borrowed_books": loan}}
	_, err := r.Session.Database(databaseName).Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ReturnBook -
func (r *Repository) ReturnBook(ctx context.Context, filter interface{}, loan domain.Loan) (interface{}, error) {

	if err := loan.Valid(); err != nil {
		return nil, err
	}

	update := bson.M{"$pop": bson.M{"lent_books": loan}}
	_, err := r.Session.Database(databaseName).Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
