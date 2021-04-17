package model

import "gorm.io/gorm"

type (
	// Profile はユーザーの情報を格納
	Profile struct {
		gorm.Model    `json:"-"`
		Name          string         `json:"Name" validate:"required"`
		Age           int            `json:"Age" validate:"required"`
		Gender        string         `json:"Gender" validate:"required"`
		FavoriteFoods []FavoriteFood `json:"FavoriteFoods" validate:"required"`
	}
)
