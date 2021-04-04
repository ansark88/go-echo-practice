package profile

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	ProfilesController struct{}

	// Profile はユーザーの情報を格納
	Profile struct {
		Name          string   `json:"Name" validate:"required"`
		Age           int      `json:"Age" validate:"required"`
		Gender        string   `json:"Gender" validate:"required"`
		FavoriteFoods []string `json:"FavoriteFoods" validate:"required"`
	}
)

// Profiles はProfile型のMap
var Profiles map[string]*Profile

// InitProfile はユーザーのプロファイルの初期データセット
func (controller ProfilesController) InitProfile() {
	Profiles = map[string]*Profile{}
	Profiles["Alice"] = &Profile{
		Name:          "Bob",
		Age:           25,
		Gender:        "Man",
		FavoriteFoods: []string{"Hamburger", "Cookie", "Chocolate"},
	}
	Profiles["Bob"] = &Profile{
		Name:          "Alice",
		Age:           24,
		Gender:        "Woman",
		FavoriteFoods: []string{"Apple", "Orange", "Melon"},
	}
}

// GetProfile はユーザーのプロフィール取得API
func (controller ProfilesController) GetProfile(ctx echo.Context) error {
	// Todo: ユーザー情報はDBに永続化すること
	name := ctx.Param("name")

	if profile, ok := Profiles[name]; ok {
		return ctx.JSON(http.StatusOK, profile)
	}

	return ctx.JSON(http.StatusBadRequest, "存在しないユーザーです")
}

// AddProfile はユーザーのプロファイル保存API
func (controller ProfilesController) AddProfile(ctx echo.Context) error {
	p := new(Profile)
	if err := ctx.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if _, ok := Profiles[p.Name]; ok {
		return echo.NewHTTPError(http.StatusBadRequest, "既に存在する名前です")
	}

	Profiles[p.Name] = &Profile{
		Name:          p.Name,
		Age:           p.Age,
		Gender:        p.Gender,
		FavoriteFoods: p.FavoriteFoods,
	}

	return ctx.NoContent(http.StatusCreated)
}
