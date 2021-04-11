package model

type (
	// Profile はユーザーの情報を格納
	Profile struct {
		Name          string   `json:"Name" validate:"required"`
		Age           int      `json:"Age" validate:"required"`
		Gender        string   `json:"Gender" validate:"required"`
		FavoriteFoods []string `json:"FavoriteFoods" validate:"required"`
	}
)
