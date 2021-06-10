package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RenderingHandler HTMLレンダリング用のHandlerのインターフェイス
type RenderingHandler interface {
	HandleMainRendering(ctx echo.Context) (err error)
}

type renderingHandler struct{}

type message struct {
	Title string
	Text  string
}

// NewRenderingHandler はコンストラクタ
func NewRenderingHandler() RenderingHandler {
	return &renderingHandler{}
}

func (rh renderingHandler) HandleMainRendering(ctx echo.Context) (err error) {
	m := &message{
		Title: "Hello World",
		Text:  "Hello!",
	}

	return ctx.Render(http.StatusOK, "profile", m)
}
