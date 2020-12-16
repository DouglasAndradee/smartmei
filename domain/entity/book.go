package entity

import "time"

// BookFetch -
type BookFetch interface {
}

// Book - It's a model of book in the system
type Book struct {
	Title    string    `json:"title"`
	Pages    uint64    `json:"pages"`
	Owner    int64     `json:"owner"`
	Loaned   bool      `json:"loaned"`
	CreateAt time.Time `json:"create_at"`
}
