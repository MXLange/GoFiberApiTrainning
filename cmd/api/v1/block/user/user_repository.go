package user

import (
	"errors"

	db "github.com/api/interface/internal/database"
	"gorm.io/gorm"
)

func (u *User) GetUserByEmailRepo() error {

	result := db.Conn.Where("email = ?", u.Email).First(&u)

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return nil
}

func (u *User) CreateUserRepo() error {

	result := db.Conn.Create(&u)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
