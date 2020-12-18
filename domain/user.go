package domain

import (
	"errors"
	"time"

	"github.com/douglasandradeee/smartmei/helper"
)

// User - It's a user model
type User struct {
	ID         interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string      `json:"name" bson:"name"`
	Email      string      `json:"email" bson:"email"`
	Collection []Book      `json:"collection" bson:"collection"`
	Lent       []Loan      `json:"lent_books" bson:"lent_books"`
	Borrowed   []Loan      `json:"borrowed_books" bson:"borrowed_books"`
	CreateAt   time.Time   `json:"created_at" bson:"created_at"`
}

// Valid - Validadtes user fields
func (u *User) Valid() error {

	if u.Name == "" {
		return errors.New("The name isn't valid")
	}

	if u.Email == "" || !u.ValidEmail() {
		return errors.New("The email isn't valid")
	}

	return nil
}

// DefaultFields - Validate non-null default fields
func (u *User) DefaultFields() {
	u.CreateAt = time.Now()
	u.Collection = []Book{}
	u.Lent = []Loan{}
	u.Borrowed = []Loan{}
}

// FindBookInCollection - Finds a book in the collection
func (u *User) FindBookInCollection(id int64) bool {
	for _, book := range u.Collection {
		if book.ID == id {
			return true
		}
	}
	return false
}

// FindBookInLent - Find a borrowed book
func (u *User) FindBookInLent(id int64) (*Loan, bool) {
	for _, loan := range u.Lent {
		if loan.BookID == id {
			return &loan, true
		}
	}
	return nil, false
}

// ValidEmail - Validates an email
func (u *User) ValidEmail() bool {
	return helper.ValidEmail(u.Email)
}
