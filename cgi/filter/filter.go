package filter

import "iogo/cgi/http"

type Filter func(ctx *http.Context) error

