package handler

import (
	"go-echo-practice/domain/model"
	"go-echo-practice/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ProfileHandler :ProfileのHandlerのインターフェイス
type ProfileHandler interface {
	HandleGetProfile(ctx echo.Context) (err error)
	HandleAddProfile(ctx echo.Context) (err error)
}

type profileHandler struct {
	profileUseCase usecase.ProfileUseCase
}

// NewProfileHandler はコンストラクタ
func NewProfileHandler(pu usecase.ProfileUseCase) ProfileHandler {
	return &profileHandler{
		profileUseCase: pu,
	}
}

func (ph profileHandler) HandleGetProfile(ctx echo.Context) (err error) {
	// Todo: ユーザー情報はDBに永続化すること
	name := ctx.Param("name")

	profile, err := ph.profileUseCase.GetProfile(name)

	// 取り敢えず400を返すが、エラーメッセージを見てDBエラーなどの時は500を返すほうが望ましい
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, profile)
}

func (ph profileHandler) HandleAddProfile(ctx echo.Context) (err error) {
	p := new(model.Profile)
	if err := ctx.Bind(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = ph.profileUseCase.AddProfile(p)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusCreated)
}
