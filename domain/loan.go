package domain

import (
	"errors"
	"time"
)

type Loan struct {
	BookID     interface{} `json:"id" bson:"id"`
	From       interface{} `json:"from_user" bson:"from_user"`
	To         interface{} `json:"to_user" bson:"to_user"`
	LentAt     time.Time   `json:"lent_at" bson:"lent_at"`
	ReturnedAt time.Time   `json:"returned_at" bson:"returned_at"`
}

func (l *Loan) Valid() error {
	if l.BookID == 0 {
		return errors.New("")
	}

	if l.From == 0 {
		return errors.New("")
	}

	if l.To == 0 {
		return errors.New("")
	}

	if l.LentAt.After(time.Now()) {
		return errors.New("")
	}
	return nil
}
