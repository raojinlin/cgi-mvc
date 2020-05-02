package router

import (
	"iogo/cgi/http"
	"testing"
)

func TestRouter_Get(t *testing.T) {
	router := NewRouter("/")
	router.Get("/", func(ctx *http.Context) (i int, s string) {
		return 0, ""
	})
}
