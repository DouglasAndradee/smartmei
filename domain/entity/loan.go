package entity

import "time"

// LoanFetch -
type LoanFetch interface {
}

// Loan - It's a model of loan in the system
type Loan struct {
	From     int64       `json:"from"`
	To       int64       `json:"to"`
	BookID   interface{} `json:"book"`
	CreateAt time.Time   `json:"create_at"`
}
