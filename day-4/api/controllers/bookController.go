package controllers

import (
	"net/http"
	"strconv"
	"time"

	"api/helpers"
	"api/models"

	"github.com/labstack/echo"
)

var (
	Books = map[int]*models.Book{}
	seq   = 1
)

func GetAllBook(c echo.Context) error {
	if len(Books) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": helpers.EmptyData,
			"data":    nil,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": helpers.Success,
		"data":    Books,
	})
}

func GetOneBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	helpers.PanicIfError(err)
	if Books[id] == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": helpers.BookNotFound,
			"data":    nil,
		})
	}
	res := helpers.BookResponse(*Books[id])
	return c.JSON(http.StatusOK, res)
}

func CreateBook(c echo.Context) error {
	req := &models.Book{
		ID:        seq,
		CreatedAt: time.Now().Format(time.RFC822),
		UpdatedAt: time.Now().Format(time.RFC822),
	}
	data, err := helpers.BookRequest(req, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": helpers.ErrorInput,
			"data":    nil,
		})
	}
	Books[req.ID] = data
	seq++
	res := helpers.BookResponse(*data)
	return c.JSON(http.StatusCreated, res)
}

func UpdateBook(c echo.Context) error {
	req := new(models.Book)
	id, err := strconv.Atoi(c.Param("id"))
	helpers.PanicIfError(err)
	if Books[id] == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": helpers.BookNotFound,
			"data":    nil,
		})
	}

	data, err := helpers.BookRequest(req, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": helpers.ErrorInput,
			"data":    nil,
		})
	}

	res := Books[id]
	res.Title = data.Title
	res.Isbn = data.Isbn
	res.Writer = data.Writer
	res.UpdatedAt = time.Now().Format(time.RFC822)
	resp := helpers.BookResponse(*res)
	return c.JSON(http.StatusOK, resp)
}

func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	helpers.PanicIfError(err)
	if Books[id] == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": helpers.BookNotFound,
			"data":    nil,
		})
	}
	data := Books[id]
	data.DeletedAt = time.Now().Format(time.RFC822)
	res := helpers.BookResponse(*data)
	delete(Books, id)
	return c.JSON(http.StatusOK, res)
}
