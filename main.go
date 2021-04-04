package main

import (
	"net/http"

	"go-echo-pracetice/fizzbuzz"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	api := echo.New()
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// 練習問題1
	fc := fizzbuzz.FizzBuzzController{}
	api.GET("FizzBuzz/:num", fc.FizzBuzz)

	api.Logger.Fatal(api.Start(":8080"))
}
