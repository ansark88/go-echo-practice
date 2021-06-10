package template

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template はHTMLのレンダリングのテンプレート
type Template struct {
	templates *template.Template
}

// NewTemplate はコンストラクタ
func NewTemplate() *Template {
	t := &Template{
		templates: template.Must(template.ParseGlob("template/*.tpl")),
	}
	return t
}

// Render は echoのRenderメソッドが呼び出す
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
