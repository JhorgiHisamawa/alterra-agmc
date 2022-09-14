package helpers

import "api/models"

func Response(data models.User) *models.UserResponse {
	res := &models.UserResponse{
		Message: "success",
		Data:    data,
	}
	return res
}
