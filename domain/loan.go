package domain

import (
	"errors"
	"time"
)

// Loan - It's a loan model
type Loan struct {
	BookID     int64     `json:"book_id" bson:"id"`
	From       int64     `json:"from_user" bson:"from_user"`
	To         int64     `json:"to_user" bson:"to_user"`
	LentAt     time.Time `json:"lent_at" bson:"lent_at"`
	ReturnedAt time.Time `json:"returned_at" bson:"returned_at"`
}

// Valid - Validate the book's field
func (l *Loan) Valid() error {
	if l.BookID <= 0 {
		return errors.New("BookID is not valid, because it's need to be greater or equal 0")
	}

	if l.From <= 0 {
		return errors.New("The id is not valid, because it's need to be greater or equal 0")
	}

	if l.To <= 0 {
		return errors.New("The id is not valid, because it's need to be greater or equal 0")
	}
	return nil
}
