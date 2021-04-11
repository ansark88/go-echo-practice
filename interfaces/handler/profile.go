package interfaces

import (
	"go-echo-practice/domain/model"
	"go-echo-practice/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ProfileHandler :ProfileのHandlerのインターフェイス
type ProfileHandler interface {
	HandleGetProfile(ctx echo.Context)
	HandleAddProfile(ctx echo.Context)
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

func (ph profileHandler) HandleGetProfile(ctx echo.Context) {
	// Todo: ユーザー情報はDBに永続化すること
	name := ctx.Param("name")

	profile, err := usecase.ProfileUseCase.GetProfile(name)

	if err != nil {
		response.HTTPError("http.StatusBadRequest, err.Error()")
	}

	// if profile, ok := Profiles[name]; ok {
	// 	return ctx.JSON(http.StatusOK, profile)
	// }

	return response.JSON(http.StatusBadRequest, "存在しないユーザーです")
}

func (ph profileHandler) HandleAddProfile(ctx echo.Context) {
	p := new(model.Profile)
	if err := ctx.Bind(p); err != nil {
		response.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(p); err != nil {
		response.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := usecase.ProfileUseCase.AddProfile(p)

	if err != nil {
		response.NewHTTPError(http.StatusBadRequest, err.Error())
	}
}
