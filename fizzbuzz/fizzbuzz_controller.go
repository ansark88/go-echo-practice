package fizzbuzz

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	FizzBuzzController struct{}
)

// FizzBuzz は1からnumまでの数のfizzbuzzを返す
func (controller FizzBuzzController) FizzBuzz(ctx echo.Context) error {

	num, err := strconv.Atoi(ctx.Param("num"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if num < 0 {
		return ctx.JSON(http.StatusBadRequest, "numは正の数にしてください")
	}

	var s string
	var result []string
	for i := 1; i <= num; i++ {
		if i%15 == 0 {
			s = "FizzBuzz!"
		} else if i%5 == 0 {
			s = "Buzz"
		} else if i%3 == 0 {
			s = "Fizz"
		} else {
			s = strconv.Itoa(i)
		}

		result = append(result, s)
	}

	return ctx.JSON(http.StatusOK, result)
}
