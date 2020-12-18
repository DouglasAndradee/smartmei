package domain_test

import (
	"testing"
	"time"

	"github.com/douglasandradeee/smartmei/domain"
	"github.com/stretchr/testify/assert"
)

var user = domain.User{
	ID:         1,
	Name:       "Douglas",
	Email:      "dba@hotmail.com",
	Collection: []domain.Book{},
	Lent:       []domain.Loan{},
	Borrowed:   []domain.Loan{},
	CreateAt:   time.Now(),
}

func TestUserNotEmpty(t *testing.T) {
	assert.NotEmpty(t, user, "the user is not empty")
}

func TestUserIds(t *testing.T) {
	assert.GreaterOrEqual(t, 1, user.ID, "The user id is valid")
}

func TestUserIsNotNil(t *testing.T) {
	assert.NotNil(t, loan, "The user is not nil.")
}

func TestUserListIsEmpty(t *testing.T) {
	assert.Empty(t, user.Collection, "The user collection of book is empty")
	assert.Empty(t, user.Lent, "The user lent of book is empty")
	assert.Empty(t, user.Borrowed, "The user borrowed of book is empty")
}

func TestValidEmail(t *testing.T) {
	assert.Equal(t, true, user.ValidEmail(), "The user's email is valid.")
}
