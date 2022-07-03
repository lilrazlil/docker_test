package api

import "gorm.io/gorm"

type userResponse struct {
	gorm.Model
	Name  string
	Email string
}
