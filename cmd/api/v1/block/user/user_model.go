package user

import "gorm.io/gorm"

type User struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Surename   string `json:"surename"`
	Email      string `json:"email" gorm:"unique"`
	gorm.Model `json:"-"`
}
