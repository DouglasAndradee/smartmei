package domain

import (
	"errors"
	"time"
)

type User struct {
	ID         interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string      `json:"name" bson:"name"`
	Email      string      `json:"email" bson:"email"`
	Collection []Book      `json:"collection" bson:"collection"`
	Lent       []Loan      `json:"lent_books" bson:"lent_books"`
	Borrowed   []Loan      `json:"borrowed_books" bson:"borrowed_books"`
	CreateAt   time.Time   `json:"created_at" bson:"created_at"`
}

func (u *User) Valid() error {
	if u.Name == "" || u.Email == "" {
		return errors.New("")
	}

	return nil
}

func (u *User) DefaultFields() {
	u.CreateAt = time.Now()
	u.Collection = []Book{}
	u.Lent = []Loan{}
	u.Borrowed = []Loan{}
}

func (u *User) FindBookInCollection(id int64) bool {
	for _, book := range u.Collection {
		if book.ID == id {
			return true
		}
	}
	return false
}

func (u *User) FindBookInLent(id int64) (*Loan, bool) {
	for _, loan := range u.Lent {
		if loan.BookID == id {
			return &loan, true
		}
	}
	return nil, false
}
