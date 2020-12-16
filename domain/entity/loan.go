package entity

import "time"

// LoanFetch -
type LoanFetch interface {
}

// Loan - It's a model of loan in the system
type Loan struct {
	From     interface{} `json:"from"`
	To       interface{} `json:"to"`
	BookID   interface{} `json:"book"`
	CreateAt time.Time   `json:"create_at"`
}
