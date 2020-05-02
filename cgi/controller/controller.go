package controller

import (
	"bytes"
	"fmt"
	"iogo/cgi/http"
)

type Handler func(ctx *http.Context) (int, string)

type Controller interface {
	Output()
}

type BaseController struct {
	context *http.Context
	handler Handler
}


func header(field, value string)  {
	fmt.Printf("%s: %s\r\n", field, value)
}

func (c *BaseController) Filter() error {
	return nil
}

func (c *BaseController) Middleware() error  {
	return nil
}

func (c *BaseController) Output()  {
	var err error
	if c.context.Response.GetHeader("Content-Type") == "" {
		c.context.Response.SetHeader("Content-Type", "text/html")
	}

	var status int
	var result string

	err = c.Filter()
	err = c.Middleware()

	if err != nil {
		status = 500
		result = "Server Internal Error"
	} else {
		status, result = c.handler(c.context)
	}

	if c.context.Response.Body == nil {
		var bf bytes.Buffer
		bf.WriteString(result)
		c.context.Response.Body = &bf
		c.context.Response.SetLen(bf.Len())
	}

	c.context.Response.SetStatus(status)

	for field, value := range c.context.Response.Headers {
		header(field, value)
	}

	for _, cookie := range c.context.Response.GetCookies() {
		if cookie != nil {
			header("Set-Cookie", cookie.String())
		}
	}

	fmt.Printf("\r\n")
	for  {
		line, err := c.context.Response.Body.ReadBytes('\n')

		fmt.Print(string(line))
		if err != nil {
			break
		}
	}
}



func NewController(ctx *http.Context, handler Handler) *BaseController {
	return &BaseController{context: ctx, handler: handler}
}
