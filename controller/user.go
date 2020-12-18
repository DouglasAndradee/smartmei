package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/DouglasAndradee/smartmei/domain"
	"github.com/DouglasAndradee/smartmei/repository"
	"github.com/labstack/echo/v4"
)

// InsertUser -
func InsertUser(r *repository.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := domain.User{}
		if err := c.Bind(&user); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		ctx := context.TODO()

		count, err := r.CountUser(ctx)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user.ID = *count + 1
		result, err := r.InsertUser(ctx, user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, result)
	}
}

// GetUser -
func GetUser(r *repository.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.TODO()

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		filter := make(map[string]interface{})
		filter["_id"] = id

		result, err := r.GetUser(ctx, filter)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, result)
	}
}

// InsertBookToUser -
func InsertBookToUser(r *repository.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		body := echo.Map{}
		if err := c.Bind(&body); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		id, ok := body["logged_user_id"]
		if !ok {
			return c.JSON(http.StatusBadRequest, "Not id")
		}

		title, ok := body["title"]
		if !ok {
			return c.JSON(http.StatusBadRequest, "Not Title")
		}

		pages, ok := body["pages"]
		if !ok {
			return c.JSON(http.StatusBadRequest, "Not Pages")
		}

		ctx := context.TODO()

		filter := make(map[string]interface{})
		filter["_id"] = id

		user, err := r.GetUser(ctx, filter)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		book := domain.Book{}
		book.ID = int64(len(user.Collection) + 1)
		book.Title = title.(string)

		book.Pages = pages

		book.DefaultFields()

		_, err = r.AddBook(ctx, filter, book)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, book)
	}
}

// LendBook -
func LendBook(r *repository.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		body := echo.Map{}
		if err := c.Bind(&body); err != nil {
			return c.NoContent(http.StatusOK)
		}

		fmt.Println(body)
		from_user_id, ok := body["logged_user_id"]
		if !ok {
			return c.JSON(http.StatusBadRequest, "Not id")
		}

		book_id, ok := body["book_id"]
		if !ok {
			return c.JSON(http.StatusBadRequest, "Not Title")
		}

		to_user_id, ok := body["to_user_id"]
		if !ok {
			return c.JSON(http.StatusBadRequest, "Not Pages")
		}

		ctx := context.TODO()

		filter := make(map[string]interface{})
		filter["_id"] = from_user_id

		user, err := r.GetUser(ctx, filter)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if !user.FindBookInCollection(1) {
			return c.JSON(http.StatusBadRequest, "The user hasn't book")
		}

		_, flag := user.FindBookInLent(1)
		if flag {
			return c.JSON(http.StatusBadRequest, "The book is already borrowed")
		}

		loan := domain.Loan{}
		loan.BookID = book_id
		loan.From = from_user_id
		loan.To = to_user_id
		loan.LentAt = time.Now()
		loan.ReturnedAt = time.Now().Add(time.Hour * 48)

		_, err = r.LendBook(ctx, filter, loan)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, loan)
	}
}

// ReturnBook -
func ReturnBook(r *repository.Repository) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := echo.Map{}
		if err := c.Bind(u); err != nil {
			return c.NoContent(http.StatusOK)
		}
		return c.NoContent(http.StatusOK)
	}
}
