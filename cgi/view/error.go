package view

import "iogo/cgi/http"

type ErrorView struct {
	*View
}

type ErrorData struct {
	Text string
	Title string
	Status int
}

func (errorView *ErrorView) Render(title, text string, status int) string  {
	var errorData = ErrorData{
		Title: title,
		Text: text,
		Status: status,
	}
	return errorView.View.Render(errorData)
}

func NewErrorView(ctx *http.Context) *ErrorView {
	return &ErrorView{
		View:   NewView(ctx.Config.TemplatePath + "error.html", ctx),
	}
}

