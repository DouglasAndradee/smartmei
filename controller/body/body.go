package body

import (
	"errors"

	"github.com/douglasandradeee/smartmei/helper"
)

// Lend -
type Lend struct {
	BookID int64 `json:"book_id"`
	From   int64 `json:"logged_user_id"`
	To     int64 `json:"to_user_id,omitmepty"`
}

// Valid -
func (l *Lend) Valid(destination bool) error {
	if l.BookID <= 0 {
		return errors.New("BookID is not valid, because it's need to be greater or equal 0")
	}

	if l.From <= 0 {
		return errors.New("The id is not valid, because it's need to be greater or equal 0")
	}

	if destination {
		if l.To <= 0 {
			return errors.New("The id is not valid, because it's need to be greater or equal 0")
		}
	}
	return nil
}

// User -
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ValidEmail -
func (u *User) ValidEmail() bool {

	return helper.ValidEmail(u.Email)
}

// Book -
type Book struct {
	ID    int64  `json:"logged_user_id"`
	Title string `json:"title"`
	Pages string `json:"pages"`
}

// Valid -
func (b *Book) Valid() error {

	if b.ID == 0 {
		return errors.New("The id is not valid, because it's need to be greater or equal 0")
	}

	if b.Title == "" {
		return errors.New("The title is not valid, because it empty")
	}

	return nil
}
