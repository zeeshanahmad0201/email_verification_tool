package email

import (
	"strings"

	"github.com/zeeshanahmad0201/email_verification_tool/pkg/helpers"
)

func VerifyEmail(email string) error {

	if err := helpers.ValidateEmail(email); err != nil {
		return err
	}

	domain := email[strings.Index(email, "@")+1:]
	if err := CheckDomain(domain); err != nil {
		return err
	}

	return nil
}
