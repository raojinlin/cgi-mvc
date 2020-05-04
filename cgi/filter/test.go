package filter

import (
	"bytes"
	"fmt"
	"iogo/cgi/http"
	"iogo/cgi/view"
	"strings"
)

var AdminPathFilter Filter = func(ctx *http.Context) error {
	if !strings.Contains(ctx.Request.Uri, "admin") {
		return nil
	}

	var err = fmt.Errorf("error")
	var bf bytes.Buffer
	bf.WriteString(view.NewErrorView(ctx).Render(
		"Admin access denied",
		"Administrator zone, Access denied.",
		403,
	))

	ctx.Response.SetStatus(403)
	ctx.Response.Body = &bf

	return err
}

var InvalidQueryParamsFilter Filter = func(ctx *http.Context) error {
	if sql, ok := ctx.Request.QueryParams["sql"]; !ok || sql == "" {
		return nil
	}

	var err = fmt.Errorf("Invalid query params \"sql\"")
	var bf bytes.Buffer
	bf.WriteString(view.NewErrorView(ctx).Render(err.Error(), err.Error(), 400))

	ctx.Response.SetStatus(400)
	ctx.Response.Body = &bf

	return err
}