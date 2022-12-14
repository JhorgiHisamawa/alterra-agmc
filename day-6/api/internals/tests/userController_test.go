package controllers

import (
	s "api/database/seeder"
	"api/internals/config"
	"api/internals/controllers"
	"api/internals/models"
	"api/pkg/util"
	"fmt"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func DBInit(t *testing.T) {
	config.DbManager()
	config.ConnectDB()
	s.DeleteSeeder()
	err := s.NewUserSeeder()
	util.PanicIfError(err)
}

func DBInitEmpty(t *testing.T) {
	config.DbManager()
	config.ConnectDB()
	s.DeleteSeeder()
}

var token string

func TestLogin(t *testing.T) {
	DBInit(t)
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	body := models.User{
		Email:    "testing@gmail.com",
		Password: "passwordtest",
	}

	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, controllers.Login(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	dataLogin := result["data"].(map[string]interface{})
	assert.Equal(t, util.Success, result["message"])
	token = dataLogin["token"].(string)
}
func TestGetAllUserEmpty(t *testing.T) {
	DBInitEmpty(t)
	//setup echo context
	e := echo.New()
	//setup request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, controllers.GetAllUser(c))
	assert.Equal(t, http.StatusNotFound, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, util.EmptyData, result["message"])
}

func TestCreateUserSuccess(t *testing.T) {
	//setup echo context
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	body := models.User{
		Name:     "testing",
		Email:    "testing@gmail.com",
		Password: "passwordtest",
		Role:     "test",
	}

	//setup request
	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, controllers.CreateUser(c))
	assert.Equal(t, http.StatusCreated, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, util.Success, result["message"])
}

func TestCreateUserBadRequest(t *testing.T) {
	//setup echo context
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	body := models.User{
		Name:     "testing",
		Email:    "",
		Password: "passwordtest",
		Role:     "test",
	}

	//setup request
	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, controllers.CreateUser(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, util.ErrorInput, result["message"])
}

func TestGetAllUsersSuccess(t *testing.T) {
	DBInit(t)
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, controllers.GetAllUser(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, util.Success, result["message"])
}

func TestGetOneUserNotFound(t *testing.T) {
	DBInit(t)
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("9")

	//test
	assert.NoError(t, controllers.GetOneUser(c))
	assert.Equal(t, http.StatusNotFound, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, util.UserNotFound, result["message"])
}

func TestGetOneUserSuccess(t *testing.T) {
	DBInit(t)
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, controllers.GetOneUser(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, util.Success, result["message"])
}

func TestUpdateUserButUserNotFound(t *testing.T) {
	DBInit(t)
	//setup echo context
	e := echo.New()

	body := models.User{
		Name:     "testing",
		Email:    "",
		Password: "passwordtest",
		Role:     "test",
	}

	//setup request
	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(string(data)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("9")

	//test
	assert.NoError(t, controllers.UpdateUser(c))
	assert.Equal(t, http.StatusNotFound, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, util.UserNotFound, result["message"])
}

func TestUpdateUserBadRequest(t *testing.T) {
	DBInit(t)
	//setup echo context
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	body := models.User{
		Name:     "testing",
		Email:    "testing@gmail.com",
		Password: "passwordtest",
		Role:     "test",
	}

	//setup request
	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	fmt.Println(token, "ini token slur")
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, controllers.UpdateUser(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, util.ErrorInput, result["message"])
}

func TestUpdateUserSucccess(t *testing.T) {
	DBInit(t)
	//setup echo context
	e := echo.New()
	e.Validator = &config.CustomValidator{Validator: validator.New()}

	body := models.User{
		Name:     "testing",
		Email:    "testingupdate@gmail.com",
		Password: "passwordtest",
		Role:     "test",
	}

	//setup request
	data, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(string(data)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, controllers.UpdateUser(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, util.Success, result["message"])
}

func TestDeleteUserButBookNotFound(t *testing.T) {
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("9")

	//test
	assert.NoError(t, controllers.DeleteUser(c))
	assert.Equal(t, http.StatusNotFound, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, util.UserNotFound, result["message"])
}

func TestDeleteUserSuccess(t *testing.T) {
	DBInit(t)
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, controllers.DeleteUser(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, util.Success, result["message"])
}
