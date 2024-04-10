package user

import (
	"errors"
)

func (u *User) ValidateNewUser() ([]string, error) {
	var err error = nil
	errArray := []string{}

	if u.Name == "" {
		errArray = append(errArray, "Name is required")
	}
	if u.Surename == "" {
		errArray = append(errArray, "Surename is required")
	}
	if u.Email == "" {
		errArray = append(errArray, "Email is required")
	}
	if len(errArray) > 0 {
		err = errors.New("validation error")
	}
	return errArray, err
}
