package view

import "iogo/cgi/http"

type IndexView struct {
	*View
}

type IndexData struct {
	Name string
}

func NewIndexView(ctx *http.Context) *IndexView {
	return &IndexView{NewView(ctx.Config.TemplatePath + "index.html", ctx)}
}