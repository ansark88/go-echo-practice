package main

import (
	"log"
	"net/http"
	"os"

	"go-echo-practice/database"
	"go-echo-practice/fizzbuzz"
	"go-echo-practice/interfaces/handler"
	"go-echo-practice/template"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
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
	// envファイル読み込み
	err := godotenv.Load(".env.sample")
	if err != nil {
		log.Fatal("envファイル読み込みエラー")
	}

	api := echo.New()
	api.Validator = &CustomValidator{validator: validator.New()}
	api.Renderer = template.NewTemplate()
	api.Static("/static", "assets")
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	// DB設定
	db := database.GetInstance()
	err = database.Migrate(db)
	database.SetInitData(db)

	if err != nil {
		log.Fatal("DB起動エラー")
		os.Exit(1)
	}

	// DIによる初期設定
	profileHandler := profileInitialize()

	// 練習問題6
	renderingHandler := handler.NewRenderingHandler()
	api.GET("/", renderingHandler.HandleMainRendering)

	// 練習問題1
	fc := fizzbuzz.FizzBuzzController{}
	api.GET("FizzBuzz/:num", fc.FizzBuzz)

	// 練習問題2,3,4
	api.GET("Profile", profileHandler.HandleListProfile)
	api.GET("Profile/:name", profileHandler.HandleGetProfile)
	api.POST("Profile", profileHandler.HandleAddProfile)

	api.Logger.Fatal(api.Start(":8080"))
}
