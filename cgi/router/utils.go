package router

import (
	"regexp"
	"strings"
)

func routeCompare(route, target string) bool {
	return string2regexp(route).MatchString(target)
}

func string2regexp(route string) *regexp.Regexp {
	paramsReg := regexp.MustCompile(":(?P<param>\\w+)(?P<optional>\\?)?")
	substitution := "(?m)(?P<$param>[^\\/]+)$optional"

	regexpStr := paramsReg.ReplaceAllString(route, substitution)
	if !strings.HasPrefix(regexpStr, "$") {
		regexpStr += "[/]{0,}$"
	}

	return regexp.MustCompile(regexpStr)
}

func getRouteParams(route, target string) map[string]string {
	reg := string2regexp(route)
	match := reg.FindStringSubmatch(target)

	params := make(map[string]string)
	for i, name := range reg.SubexpNames() {
		if i > 0 && i <= len(match) {
			params[name] = match[i]
		}
	}


	return params
}

func getRouteKey(method, path, root string) string {
	if strings.HasSuffix(root, "/") {
		root = root[:len(root) - 1]
	}

	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	if path == "/" {
		path = ""
	}

	var p string
	if path == "/" || path == "" {
		p = root
	} else {
		p = root + "/" + path
	}

	return strings.ToUpper(method) + "#" + p
}

