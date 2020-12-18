package repository

import (
	"context"
	"time"

	"github.com/douglasandradeee/smartmei/domain"
	"go.mongodb.org/mongo-driver/bson"
)

// NewUser -
func (r *Repository) NewUser(id int64, name, email string) domain.User {
	user := domain.User{ID: id, Name: name, Email: email}
	user.DefaultFields()
	return user
}

// NewBook -
func (r *Repository) NewBook(id int64, title string, pages string) domain.Book {
	book := domain.Book{ID: id, Title: title, Pages: pages}
	book.DefaultFields()
	return book
}

// NewLoan -
func (r *Repository) NewLoan(id, from, to int64) domain.Loan {
	loan := domain.Loan{BookID: id, From: from, To: to}
	loan.LentAt = time.Now()
	loan.ReturnedAt = time.Now().Add(time.Hour * 48)
	return loan
}

// InsertUser - Insert a new user in the database
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

// CountUser - Assign user id
func (r *Repository) CountUser(ctx context.Context) (*int64, error) {
	filter := bson.M{}
	count, err := r.Session.Database(databaseName).Collection("users").CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	return &count, nil
}

// GetUser - Get user in the database
func (r *Repository) GetUser(ctx context.Context, filter interface{}) (*domain.User, error) {
	result := domain.User{}
	err := r.Session.Database(databaseName).Collection("users").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// FoundBook - Finds a book in the database
func (r *Repository) FoundBook(ctx context.Context, filter interface{}) (interface{}, error) {
	result := domain.User{}
	err := r.Session.Database(databaseName).Collection("users").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// AddBook - Assigns a book to a user
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

// LendBook - Lend a book another user
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

// ReturnBook - Return de borrowed book
func (r *Repository) ReturnBook(ctx context.Context, filter interface{}, loan domain.Loan) (interface{}, error) {

	if err := loan.Valid(); err != nil {
		return nil, err
	}

	update := bson.M{"$pull": bson.M{"lent_books": loan}}
	_, err := r.Session.Database(databaseName).Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
