package persistence

import (
	"go-echo-practice/database"
	"go-echo-practice/domain/model"
	"go-echo-practice/domain/repository"
)

var profiles map[string]*model.Profile

type profilePersistence struct{} //何も持たない

// NewProfilePersistence は ProfileRepositoryインターフェイスを実装している
func NewProfilePersistence() repository.ProfileRepository {
	profiles = make(map[string]*model.Profile)
	return &profilePersistence{}
}

// リポジトリの実装
func (pp profilePersistence) GetProfile(name string) (profile *model.Profile, err error) {
	db := database.GetInstance()
	err = db.First(&profile, "Name = ?", name).Error

	if err != nil {
		return nil, nil // 該当なし
	}

	// リレーション先を取得する
	db.Model(&profile).Association("FavoriteFoods").Find(&profile.FavoriteFoods)

	return profile, nil
}

func (pp profilePersistence) AddProfile(p *model.Profile) error {
	db := database.GetInstance()
	err := db.Create(&p).Error

	return err
}
