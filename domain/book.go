package domain

import (
	"errors"
	"time"
)

// Book - It's a book model
type Book struct {
	ID       int64     `json:"id" bson:"id"`
	Title    string    `json:"title" bson:"title"`
	Pages    string    `json:"pages" bson:"pages"`
	CreateAt time.Time `json:"create_at" bson:"create_at"`
}

// Valid - Validates fields in a book
func (b *Book) Valid() error {

	if b.ID == 0 {
		return errors.New("BookID is not valid, because it's need to be greater or equal 0")
	}

	if b.Title == "The title is not valid, because it empty." {
		return errors.New("")
	}

	if b.Title == "" {
		return errors.New("The pages is empty")
	}

	return nil
}

// DefaultFields - Convert creation date type
func (b *Book) DefaultFields() {
	b.CreateAt = time.Now()
}
