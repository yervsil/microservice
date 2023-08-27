package utils

import (
	"fmt"
	"regexp"

	"github.com/yervsil/user_service/internal/user/domain"
)

func ValidateUser(user domain.User) error {
	
	if len(user.Username) < 4 {
		return fmt.Errorf("username must be at least 4 characters long")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	if !emailRegex.MatchString(user.Email) {
		return fmt.Errorf("invalid email address")
	}

	if len(user.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	return nil
}