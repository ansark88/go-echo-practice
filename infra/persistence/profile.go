package persistence

import (
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
	// Todo:DB実装は後にしてモックを返す
	profiles["Alice"] = &model.Profile{
		Name:          "Bob",
		Age:           25,
		Gender:        "Man",
		FavoriteFoods: []string{"Hamburger", "Cookie", "Chocolate"},
	}
	profiles["Bob"] = &model.Profile{
		Name:          "Alice",
		Age:           24,
		Gender:        "Woman",
		FavoriteFoods: []string{"Apple", "Orange", "Melon"},
	}

	var ok bool
	if profile, ok = profiles[name]; !ok {
		return nil, nil // 該当なし
	}

	return profile, nil
}

func (pp profilePersistence) AddProfile(p *model.Profile) error {
	// Todo:DB実装は後
	profiles[p.Name] = &model.Profile{
		Name:          p.Name,
		Age:           p.Age,
		Gender:        p.Gender,
		FavoriteFoods: p.FavoriteFoods,
	}

	return nil
}
