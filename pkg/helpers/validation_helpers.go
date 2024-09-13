package helpers

import (
	"fmt"
	"net/mail"
	"regexp"
)

func ValidateEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email address is required")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("invalid email address")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if ok, _ := regexp.MatchString(emailRegex, email); !ok {
		return fmt.Errorf("invalid email address")
	}

	return nil
}
