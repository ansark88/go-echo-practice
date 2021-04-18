package database

import (
	"go-echo-practice/domain/model"

	"gorm.io/gorm"
)

// Migrate はGormのマイグレーションを行う
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.Profile{}, &model.FavoriteFood{})
}
