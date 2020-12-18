package body_test

import (
	"testing"

	"github.com/douglasandradeee/smartmei/controller/body"
	"github.com/stretchr/testify/assert"
)

var book = body.Book{
	ID:    1,
	Title: "Pequeno Principe",
	Pages: "340",
}

func TestBodyBookIsNotEmpty(t *testing.T) {
	assert.NotEmpty(t, book, "The book is not empty.")
}

func TestBodyBookIsNotNil(t *testing.T) {
	assert.NotNil(t, book, "The book is not nil.")
}

var lend = body.Lend{
	BookID: 1,
	From:   1,
	To:     2,
}

func TestBodyLendIsNotEmpty(t *testing.T) {
	assert.NotEmpty(t, lend, "The lend is not empty.")
}

func TestBodyLendIsNotNil(t *testing.T) {
	assert.NotNil(t, book, "The lend is not nil.")
}

func TestBodyLendIdIsValid(t *testing.T) {
	assert.GreaterOrEqual(t, int64(1), lend.BookID, "The lend's body book_id is valid")
	assert.GreaterOrEqual(t, int64(1), lend.From, "The lend's body from_user_id is valid")
	assert.GreaterOrEqual(t, int64(2), lend.To, "The lend's body to_user_id is valid")
}

var user = body.User{
	Name:  "Douglas Borges",
	Email: "dba@hmail.com",
}

func TestBodyUserName(t *testing.T) {
	assert.Equal(t, "Douglas Borges", user.Name, "The user's email is valid.")
	assert.NotNil(t, "Douglas Borges", user.Name, "The user's email is valid.")
	assert.NotEmpty(t, "Douglas Borges", user.Name, "The user's email is valid.")
}

func TestBodyUserEmail(t *testing.T) {
	assert.Equal(t, true, user.ValidEmail(), "The user's email is valid.")
}
