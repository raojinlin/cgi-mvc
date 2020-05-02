package view

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"iogo/cgi/http"
	"os"
)

type View struct {
	Template string
	ctx *http.Context
}

func (v *View) getContent() string {
	f, err := os.OpenFile(v.Template, os.O_RDONLY, 0)
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(f)
	if err !=nil {
		panic(err)
	}

	return string(data)
}

func (v *View) Render(data interface{}) string  {
	tpl, err := template.New("index").Parse(v.getContent())
	if err != nil {
		return ""
	}

	var writer bytes.Buffer
	err = tpl.Execute(&writer, data)

	return writer.String()
}

func NewView(template string, ctx *http.Context) *View {
	return &View{Template: template, ctx: ctx}
}
