package users

import (
	"fmt"
	"log"

	"github.com/mxygem/cobrainitless/internal/accounts"
)

func Create(api *accounts.Client, email string) error {
	if email == "" {
		return fmt.Errorf("cannot create user, no email provided")
	}
	log.Printf("creating user with email: %q\n", email)

	return nil
}
