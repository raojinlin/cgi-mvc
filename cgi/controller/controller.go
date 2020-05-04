package controller

import (
	"bytes"
	"fmt"
	"iogo/cgi/filter"
	"iogo/cgi/http"
)

type Handler func(ctx *http.Context) (int, string)

type Controller interface {
	Output()
}

type BaseController struct {
	context *http.Context
	handler Handler
	filters []*filter.Filter
}


func header(field, value string)  {
	fmt.Printf("%s: %s\r\n", field, value)
}

func (c *BaseController) filter() error {
	var err error
	for _, f := range c.filters {
		err = (*f)(c.context)
		if err != nil {
			return err
		}
	}
	return err
}

func (c *BaseController) AddFilter(filter *filter.Filter) {
	c.filters = append(c.filters, filter)
}

func (c *BaseController) output()  {
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


func (c *BaseController) Output()  {
	var err error
	if c.context.Response.GetContentType() == "" {
		c.context.Response.SetContentType("text/html")
	}

	err = c.filter()
	if err != nil {
		c.output()
		return
	}

	status, result := c.handler(c.context)
	if c.context.Response.Body == nil {
		var bf bytes.Buffer
		bf.WriteString(result)
		c.context.Response.Body = &bf
		c.context.Response.SetLen(bf.Len())
	}

	c.context.Response.SetStatus(status)
	c.output()
}

func (c *BaseController) Init() {
	c.AddFilter(&filter.AdminPathFilter)
	c.AddFilter(&filter.InvalidQueryParamsFilter)
}



func NewController(ctx *http.Context, handler Handler) *BaseController {
	c := &BaseController{
		context:    ctx,
		handler:    handler,
		filters:    make([]*filter.Filter, 0),
	}

	c.Init()
	return c
}
