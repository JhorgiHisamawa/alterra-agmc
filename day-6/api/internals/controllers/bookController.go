package controllers

import (
	"net/http"
	"strconv"
	"time"

	"api/internals/models"
	"api/pkg/util"

	"github.com/labstack/echo"
)

var (
	Books = map[int]*models.Book{}
	seq   = 1
)

func GetAllBook(c echo.Context) error {
	if len(Books) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": util.EmptyData,
			"data":    nil,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": util.Success,
		"data":    Books,
	})
}

func GetOneBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	util.PanicIfError(err)
	if Books[id] == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": util.BookNotFound,
			"data":    nil,
		})
	}
	res := util.BookResponse(*Books[id])
	return c.JSON(http.StatusOK, res)
}

func CreateBook(c echo.Context) error {
	req := &models.Book{
		ID:        seq,
		CreatedAt: time.Now().Format(time.RFC822),
		UpdatedAt: time.Now().Format(time.RFC822),
	}
	data, err := util.BookRequest(req, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": util.ErrorInput,
			"data":    nil,
		})
	}
	Books[req.ID] = data
	seq++
	res := util.BookResponse(*data)
	return c.JSON(http.StatusCreated, res)
}

func UpdateBook(c echo.Context) error {
	req := new(models.Book)
	id, err := strconv.Atoi(c.Param("id"))
	util.PanicIfError(err)
	if Books[id] == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": util.BookNotFound,
			"data":    nil,
		})
	}

	data, err := util.BookRequest(req, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": util.ErrorInput,
			"data":    nil,
		})
	}

	res := Books[id]
	res.Title = data.Title
	res.Isbn = data.Isbn
	res.Writer = data.Writer
	res.UpdatedAt = time.Now().Format(time.RFC822)
	resp := util.BookResponse(*res)
	return c.JSON(http.StatusOK, resp)
}

func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	util.PanicIfError(err)
	if Books[id] == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": util.BookNotFound,
			"data":    nil,
		})
	}
	data := Books[id]
	data.DeletedAt = time.Now().Format(time.RFC822)
	res := util.BookResponse(*data)
	delete(Books, id)
	return c.JSON(http.StatusOK, res)
}
