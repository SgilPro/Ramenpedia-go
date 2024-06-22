package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name     string
	Email    string
	Birthday string
	Token    string
}
