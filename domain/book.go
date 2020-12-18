package domain

import (
	"errors"
	"time"
)

type Book struct {
	ID       int64       `json:"id" bson:"id"`
	Title    string      `json:"title" bson:"title"`
	Pages    interface{} `json:"pages" bson:"pages"`
	CreateAt time.Time   `json:"create_at" bson:"create_at"`
}

func (b *Book) Valid() error {

	if b.ID == 0 {
		return errors.New("")
	}

	if b.Title == "" {
		return errors.New("")
	}

	return nil
}

func (b *Book) DefaultFields() {
	b.CreateAt = time.Now()
}
