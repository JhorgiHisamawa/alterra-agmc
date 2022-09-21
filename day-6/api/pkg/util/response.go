package util

import "api/internals/models"

func UserResponse(data models.User) *models.UserResponse {
	res := &models.UserResponse{
		Message: Success,
		Data:    data,
	}
	return res
}

func BookResponse(data models.Book) *models.BooksResponse {
	res := &models.BooksResponse{
		Message: Success,
		Data:    data,
	}
	return res
}
