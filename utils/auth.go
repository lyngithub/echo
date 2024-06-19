package utils

import (
	"echo/cons"
	"echo/models"
)

func IsAdmin(user *models.User) bool {
	return cons.USER_ADMIN == user.UserType
}
