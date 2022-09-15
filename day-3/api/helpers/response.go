package helpers

import "api/models"

func UserResponse(data models.User) *models.UserResponse {
	res := &models.UserResponse{
		Message: "success",
		Data:    data,
	}
	return res
}

func BookResponse(data models.Books) *models.BooksResponse {
	res := &models.BooksResponse{
		Message: "success",
		Data:    data,
	}
	return res
}
