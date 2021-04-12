package main

import (
	"net/http"

	"go-echo-practice/fizzbuzz"
	"go-echo-practice/infra/persistence"
	"go-echo-practice/interfaces/handler"
	"go-echo-practice/usecase"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

// Validate は CustomValidator型に実装する必要があるメソッド
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	api := echo.New()
	api.Validator = &CustomValidator{validator: validator.New()}
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	// 初期設定
	profilePersistence := persistence.NewProfilePersistence()
	profileUseCase := usecase.NewProfileUseCase(profilePersistence)
	profileHandler := handler.NewProfileHandler(profileUseCase)

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// 練習問題1
	fc := fizzbuzz.FizzBuzzController{}
	api.GET("FizzBuzz/:num", fc.FizzBuzz)

	api.GET("Profile/:name", profileHandler.HandleGetProfile)
	api.POST("Profile", profileHandler.HandleAddProfile)

	api.Logger.Fatal(api.Start(":8080"))
}
