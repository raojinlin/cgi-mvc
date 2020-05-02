package http

import (
	"strconv"
	"strings"
)

// normalize header filed
// example: HTTP_USER_AGENT => Http-User-Agent
//          http-user-agent => Http-User-Agent
func normalizeHeaderField(field string) string {
	normalizeField := ""

	capitalizeNext := true
	for i := 0 ; i < len(field); i++ {
		if capitalizeNext {
			normalizeField += strings.ToUpper(string(field[i]))
			capitalizeNext = false
		} else if field[i] == '_' || field[i] == '-' {
			normalizeField += "-"
			capitalizeNext = true
		} else {
			normalizeField += strings.ToLower(string(field[i]))
		}
	}

	return normalizeField
}

func getHeadersFromEnvMap(envMap map[string]string, scheme string) map[string]string {
	headers := make(map[string]string)

	for field, value := range envMap {
		if strings.HasPrefix(field, strings.ToUpper(scheme) + "_") {
			headerFields := strings.Split(field, "_")
			if len(headerFields) == 2 {
				headers[normalizeHeaderField(headerFields[1])] = value
			} else if len(headerFields) > 2 {
				key := normalizeHeaderField(strings.Join(headerFields[1:], "_"))
				headers[key] = value
			}
		}
	}

	return headers
}

func getHttpHeadersFromEnvMap(envMap map[string]string) map[string]string  {
	return getHeadersFromEnvMap(envMap, "HTTP")
}

func getEnvValue(env map[string]string, key string) string {
	if val, ok := env[key]; ok {
		return val
	}

	return ""
}

func parseCookies(cookie string) []*Cookie  {
	cookieStringArray := strings.Split(cookie, "; ")
	cookies := make([]*Cookie, len(cookieStringArray))

	for i, c := range cookieStringArray {
		keyValue :=strings.Split(c, "=")
		val := ""
		if len(keyValue) == 2 {
			val = keyValue[1]
		} else if len(keyValue) > 2 {
			val = strings.Join(keyValue[1:], "=")
		}

		cookies[i] = &Cookie{
			Key:      keyValue[0],
			Value:    val,
			Domain:   "",
			Path:     "",
			Expires:  nil,
			MaxAge:   0,
			HttpOnly: false,
			Secure:   false,
		}
	}

	return cookies
}

func parseQueryParamsFromString(queryString string) map[string]string {
	params := make(map[string]string)
	queryArray := strings.Split(queryString, "&")

	for i := 0; i < len(queryArray); i++ {
		entry := strings.Split(queryArray[i], "=")

		key := entry[0]
		value := ""
		if len(entry) == 2 {
			value = entry[1]
		} else if len(entry) > 2 {
			value = strings.Join(entry[1:], "=")
		}

		params[key] = value
	}

	return params
}

func parseInt(s string) int {
	var i int
	i64, err := strconv.ParseInt(s, 0, 0)

	if err != nil {
		i = 0
	} else {
		i = int(i64)
	}

	return i
}
