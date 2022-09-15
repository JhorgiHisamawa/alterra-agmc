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
	books = map[int]*models.Books{}
	seq   = 1
)

func GetAllBook(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}

func GetOneBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	helpers.PanicIfError(err)
	return c.JSON(http.StatusOK, books[id])
}

func CreateBook(c echo.Context) error {
	data := &models.Books{
		ID:        seq,
		CreatedAt: time.Now().Format(time.RFC822),
		UpdatedAt: time.Now().Format(time.RFC822),
	}
	res, err := helpers.BookRequest(data, c)
	helpers.PanicIfError(err)
	books[data.ID] = res
	seq++
	return c.JSON(http.StatusCreated, res)
}

func UpdateBook(c echo.Context) error {
	req := new(models.Books)
	id, _ := strconv.Atoi(c.Param("id"))

	data, err := helpers.BookRequest(req, c)
	helpers.PanicIfError(err)

	res := books[id]
	res.Title = data.Title
	res.Isbn = data.Isbn
	res.Writer = data.Writer
	res.UpdatedAt = time.Now().Format(time.RFC822)
	return c.JSON(http.StatusOK, res)
}

func DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data := books[id]
	data.DeletedAt = time.Now().Format(time.RFC822)
	res := helpers.BookResponse(*data)
	delete(books, id)
	return c.JSON(http.StatusOK, res)
}
