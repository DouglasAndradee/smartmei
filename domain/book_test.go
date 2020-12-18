package domain_test

import (
	"testing"
	"time"

	"github.com/douglasandradeee/smartmei/domain"
	"github.com/stretchr/testify/assert"
)

var book = domain.Book{
	ID:       1,
	Title:    "Pequeno Principe",
	Pages:    "340",
	CreateAt: time.Now(),
}

func TestBookNotEmpty(t *testing.T) {
	assert.NotEmpty(t, book, "the book's is not empty")
}

func TestBookIds(t *testing.T) {
	assert.GreaterOrEqual(t, int64(1), book.ID, "The book's id is valid")
}

func TestBookIsNotNil(t *testing.T) {
	assert.NotNil(t, book, "The book is not nil.")
}
