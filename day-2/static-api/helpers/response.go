package helpers

import "api/models"

func Response(data models.Books) *models.BooksResponse {
	res := &models.BooksResponse{
		Message: "success",
		Data:    data,
	}
	return res
}
