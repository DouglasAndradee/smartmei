package domain_test

import (
	"testing"
	"time"

	"github.com/douglasandradeee/smartmei/domain"
	"github.com/stretchr/testify/assert"
)

var loan = domain.Loan{
	BookID:     1,
	From:       1,
	To:         2,
	LentAt:     time.Now(),
	ReturnedAt: time.Now(),
}

func TestLoanNotEmpty(t *testing.T) {
	assert.NotEmpty(t, loan, "the loan is not empty")
}

func TestLoanIds(t *testing.T) {
	assert.GreaterOrEqual(t, int64(1), loan.BookID, "The bookID is valid")
	assert.GreaterOrEqual(t, int64(1), loan.From, "The from_user_id is valid")
	assert.GreaterOrEqual(t, int64(2), loan.To, "The to_user_id is valid")
}

func TestLoanIsNotNil(t *testing.T) {
	assert.NotNil(t, loan, "The loan is not nil.")
}
