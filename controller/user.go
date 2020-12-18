package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/douglasandradeee/smartmei/controller/body"
	"github.com/douglasandradeee/smartmei/helper"
	"github.com/douglasandradeee/smartmei/repository"
	"github.com/labstack/echo/v4"
)

// InsertUser - Make a request to insert the user
func InsertUser(r *repository.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := body.User{}
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		if !user.ValidEmail() {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "email is not valid"})
		}

		ctx := context.TODO()

		count, err := r.CountUser(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		id := *count + 1
		result, err := r.InsertUser(ctx, r.NewUser(id, user.Name, user.Email))
		if err != nil {
			if helper.IsDup(err) {
				c.JSON(http.StatusInternalServerError, map[string]string{"message": "E-mail already registered"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
	}
}

// GetUser - Make a request to get a database user
func GetUser(r *repository.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.TODO()

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		filter := make(map[string]interface{})
		filter["_id"] = id

		result, err := r.GetUser(ctx, filter)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, result)
	}
}

// InserBookToUser - Make a request to assign a book to a user
func InserBookToUser(r *repository.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		b := body.Book{}
		if err := c.Bind(&b); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		ctx := context.TODO()

		filter := make(map[string]interface{})
		filter["_id"] = b.ID

		user, err := r.GetUser(ctx, filter)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		book := r.NewBook(int64(len(user.Collection)+1), b.Title, b.Pages)
		_, err = r.AddBook(ctx, filter, book)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, book)
	}
}

// LendBook - Make a request to assign a loan for a book
func LendBook(r *repository.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		b := body.Lend{}
		if err := c.Bind(&b); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		if err := b.Valid(true); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		ctx := context.TODO()

		filter := make(map[string]interface{})
		filter["_id"] = b.From

		user, err := r.GetUser(ctx, filter)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		if !user.FindBookInCollection(b.BookID) {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "The Book not found in user's collection"})
		}

		_, flag := user.FindBookInLent(b.BookID)
		if flag == true {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "This book is already on loan"})
		}

		loan := r.NewLoan(b.BookID, b.From, b.To)
		_, err = r.LendBook(ctx, filter, loan)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, loan)
	}
}

// ReturnBook - Make a request to return a borrowed book
func ReturnBook(r *repository.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		b := body.Lend{}
		if err := c.Bind(&b); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		if err := b.Valid(false); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		ctx := context.TODO()

		filter := make(map[string]interface{})
		filter["_id"] = b.From

		user, err := r.GetUser(ctx, filter)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		loan, flag := user.FindBookInLent(b.BookID)
		if flag == false {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "This book is already on loan"})
		}

		_, err = r.ReturnBook(ctx, filter, *loan)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, loan)
	}
}
