package database

import (
	"go-echo-practice/domain/model"

	"gorm.io/gorm"
)

// SetInitData はDBの初期データをセット
func SetInitData(db *gorm.DB) {
	var profile model.Profile
	db.First(&profile, "Name = ?", "Bob")

	if profile.Name != "Bob" {
		db.Create(&model.Profile{
			Name:   "Bob",
			Age:    25,
			Gender: "Man",
			FavoriteFoods: []model.FavoriteFood{
				{Food: "Hamburger"},
				{Food: "Cookie"},
				{Food: "Chocolate"},
			}})
	}

	profile = model.Profile{} // 初期化しないと上手く取れません
	db.First(&profile, "Name = ?", "Alice")

	if profile.Name != "Alice" {
		db.Create(&model.Profile{
			Name:   "Alice",
			Age:    24,
			Gender: "Woman",
			FavoriteFoods: []model.FavoriteFood{
				{Food: "Apple"},
				{Food: "Orange"},
				{Food: "Melon"},
			}})
	}
}
