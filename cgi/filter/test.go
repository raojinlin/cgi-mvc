package filter

import (
	"bytes"
	"fmt"
	"iogo/cgi/http"
	"strings"
)

var AdminPathFilter Filter = func(ctx *http.Context) error {
	if !strings.Contains(ctx.Request.Uri, "admin") {
		return nil
	}

	var err = fmt.Errorf("error")
	var bf bytes.Buffer
	bf.WriteString("<h1>Administrator zone, Access denied.</h1>")

	ctx.Response.SetStatus(403)
	ctx.Response.Body = &bf

	return err
}