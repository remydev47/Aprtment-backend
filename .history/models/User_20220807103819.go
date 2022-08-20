package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Email       string  `json:"email"`
	Password    string  `json:"passWord"`
	SocialLogin boolean `json:"socialLogin"`
}
