package controllers

import (
	"api/config"
	"api/controllers"
	"api/helpers"
	"api/models"
	"time"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetAllBooksEmpty(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, controllers.GetAllBook(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, helpers.EmptyData, result["message"])
}

func TestCreateBookSuccess(t *testing.T) {
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	body := models.Book{
		Title:  "the titles",
		Isbn:   "1-234-5678-9101112-19",
		Writer: "the writer",
	}

	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, controllers.CreateBook(c))
	assert.Equal(t, http.StatusCreated, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, helpers.Success, result["message"])
}

func TestCreateBookBadRequest(t *testing.T) {
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	body := models.Book{
		Title:  "",
		Isbn:   "1-234-5678-9101112-19",
		Writer: "the writer",
	}

	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, controllers.CreateBook(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, helpers.ErrorInput, result["message"])
}

func TestGetAllBooksSuccess(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	cg := e.NewContext(req, rec)

	assert.NoError(t, controllers.GetAllBook(cg))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, helpers.Success, result["message"])
}

func TestGetOneBookNotFound(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("9")

	assert.NoError(t, controllers.GetOneBook(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, helpers.BookNotFound, result["message"])
}

func TestGetOneBookSuccess(t *testing.T) {
	//setup echo context
	e := echo.New()

	mockBook := models.Book{
		ID:        1,
		Title:     "the titles",
		Isbn:      "1-234-5678-9101112-19",
		Writer:    "the writer",
		CreatedAt: time.Now().Format(time.RFC822),
		UpdatedAt: time.Now().Format(time.RFC822),
		DeletedAt: "",
	}

	controllers.Books = map[int]*models.Book{
		1: &mockBook,
	}

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, controllers.GetOneBook(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, helpers.Success, result["message"])
}

func TestUpdateBookButBookNotFound(t *testing.T) {
	e := echo.New()

	mockBook := models.Book{
		ID:        1,
		Title:     "the titles",
		Isbn:      "1-234-5678-9101112-19",
		Writer:    "the writer",
		CreatedAt: time.Now().Format(time.RFC822),
		UpdatedAt: time.Now().Format(time.RFC822),
		DeletedAt: "",
	}

	controllers.Books = map[int]*models.Book{
		1: &mockBook,
	}

	body := models.Book{
		Title:  "title",
		Isbn:   "1-234-5678-9101112-19",
		Writer: "the writer",
	}

	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/books", strings.NewReader(string(data)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("9")

	assert.NoError(t, controllers.UpdateBook(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, helpers.BookNotFound, result["message"])
}

func TestUpdateBookBadRequest(t *testing.T) {
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	mockBook := models.Book{
		ID:        1,
		Title:     "the titles",
		Isbn:      "1-234-5678-9101112-19",
		Writer:    "the writer",
		CreatedAt: time.Now().Format(time.RFC822),
		UpdatedAt: time.Now().Format(time.RFC822),
		DeletedAt: "",
	}

	controllers.Books = map[int]*models.Book{
		1: &mockBook,
	}

	body := models.Book{
		Title:  "",
		Isbn:   "1-234-5678-9101112-19",
		Writer: "the writer",
	}

	req := httptest.NewRequest(http.MethodPut, "/books", strings.NewReader(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, controllers.UpdateBook(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, helpers.ErrorInput, result["message"])
}

func TestUpdateBookButBookSucccess(t *testing.T) {
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	mockBook := models.Book{
		ID:        1,
		Title:     "the titles",
		Isbn:      "1-234-5678-9101112-19",
		Writer:    "the writer",
		CreatedAt: time.Now().Format(time.RFC822),
		UpdatedAt: time.Now().Format(time.RFC822),
		DeletedAt: "",
	}

	controllers.Books = map[int]*models.Book{
		1: &mockBook,
	}

	body := models.Book{
		Title:  "title",
		Isbn:   "1-234-5678-9101112-19",
		Writer: "the writer",
	}

	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/books", strings.NewReader(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, controllers.UpdateBook(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, helpers.Success, result["message"])
}

func TestDeleteBookButBookNotFound(t *testing.T) {
	e := echo.New()

	mockBook := models.Book{
		ID:        1,
		Title:     "the titles",
		Isbn:      "1-234-5678-9101112-19",
		Writer:    "the writer",
		CreatedAt: time.Now().Format(time.RFC822),
		UpdatedAt: time.Now().Format(time.RFC822),
		DeletedAt: "",
	}

	controllers.Books = map[int]*models.Book{
		1: &mockBook,
	}

	req := httptest.NewRequest(http.MethodDelete, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("9")

	assert.NoError(t, controllers.DeleteBook(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, helpers.BookNotFound, result["message"])
}

func TestDeleteBookSuccess(t *testing.T) {
	e := echo.New()

	mockBook := models.Book{
		ID:        1,
		Title:     "the titles",
		Isbn:      "1-234-5678-9101112-19",
		Writer:    "the writer",
		CreatedAt: time.Now().Format(time.RFC822),
		UpdatedAt: time.Now().Format(time.RFC822),
		DeletedAt: "",
	}

	controllers.Books = map[int]*models.Book{
		1: &mockBook,
	}

	req := httptest.NewRequest(http.MethodDelete, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, controllers.DeleteBook(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, helpers.Success, result["message"])
}
