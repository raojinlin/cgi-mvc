package router

import (
	"fmt"
	"testing"
)

func assertTrue(t *testing.T)  {

}

func TestCompare(t *testing.T)  {
	fmt.Println(routeCompare("/home/:username/spaces/.+", "/home/user/spaces/x/y"))
	fmt.Println(routeCompare("/home/user/photos/:size?", "/home/user/photos/"))
	fmt.Println(routeCompare("/home/user/photos/:size?", "/home/user/photos/"))
	fmt.Println(routeCompare("/home/user/photos/:size?", "/home/user/photos/large"))
}

func TestGetRouteParams(t *testing.T)  {
	fmt.Println(getRouteParams("/home/:username", "/home/user"))
	fmt.Println(getRouteParams("/home/:username/:page", "/home/user/1/"))
	fmt.Println(getRouteParams("/cgi-bin/main/hello/json/:username", "/cgi-bin/main/hello/json/hello"))
}