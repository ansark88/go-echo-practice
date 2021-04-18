package model

import "gorm.io/gorm"

type (
	// Profile はユーザーの情報を格納
	Profile struct {
		gorm.Model    `json:"-"`
		Name          string         `json:"Name" validate:"required" gorm:"not null;unique"`
		Age           int            `json:"Age" validate:"required" gorm:"not null"`
		Gender        string         `json:"Gender" validate:"required" gorm:"not null"`
		FavoriteFoods []FavoriteFood `json:"FavoriteFoods" validate:"required"`
	}
)
