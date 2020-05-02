package http

import "testing"

func stringAssertEqual(actual, except string, t *testing.T)  {
	if except != actual {
		t.Error("Except \"" + except + "\", but go \"" + actual + "\"")
	}
}

func inMapCheck(m map[string]string, k string, t *testing.T) {
	if _, ok := m[k]; !ok {
		t.Error("The key '" + k + "' not in map")
	}
}

func TestNormalizeHeaderField(t *testing.T) {
	stringAssertEqual(normalizeHeaderField("ACCEPT"), "Accept", t)
	stringAssertEqual(normalizeHeaderField("USER_AGENT"), "User-Agent", t)
	stringAssertEqual(normalizeHeaderField("HTTP_USER_AGENT"), "Http-User-Agent", t)
	stringAssertEqual(normalizeHeaderField("HTTP_USER-AGENT"), "Http-User-Agent", t)
	stringAssertEqual(normalizeHeaderField("HTTP-USER_AGENT"), "Http-User-Agent", t)
}

func TestParseQueryParamsFromString(t *testing.T)  {
	params := parseQueryParamsFromString("a=1&b=3&c=3&d=5&x=3==34&y&h=3")
	stringAssertEqual("1", params["a"], t)
	stringAssertEqual("3", params["b"], t)
	stringAssertEqual("5", params["d"], t)
	stringAssertEqual("3==34", params["x"], t)
	stringAssertEqual("", params["y"], t)
	stringAssertEqual("3", params["h"], t)

	if _, ok := params["y"]; !ok {
		t.Error("params should contains key 'y'")
	}
}

func TestGetHttpHeadersFromEnvMap(t *testing.T) {
	envMap := map[string]string {
		"HTTP_USER_AGENT": "Mozilla/5.0 (X11; Linux x86_64; rv:75.0) Gecko/20100101 Firefox/75.0",
		"HTTP_COOKIE": "session_id=123",
		"HTTP_CACHE_CONTROL": "max-age=0",
		"HTTP_UPGRADE_INSECURE_REQUESTS": "1",
		"REQUEST_METHOD": "GET",
	}

	headers := getHttpHeadersFromEnvMap(envMap)
	inMapCheck(headers, "User-Agent", t)
	inMapCheck(headers, "Cookie", t)
	inMapCheck(headers, "Cache-Control", t)
	inMapCheck(headers, "Upgrade-Insecure-Requests", t)
}

func TestParseInt(t *testing.T) {
	if parseInt("12") != 12 {
		t.Error("parse int error should equals 12")
	}

	if parseInt("x12") != 0 {
		t.Error("parse int error should equals 0")
	}
}

func TestParseCookies(t *testing.T) {
	cookies := parseCookies("newssss=value; session=yes=no=yes=3;or=4")
	if len(cookies) != 2 {
		t.Errorf("Except parsed cookies length be 2")
	}

	stringAssertEqual(cookies[0].Key, "newssss", t)
	stringAssertEqual(cookies[0].Value, "value", t)
	stringAssertEqual(cookies[1].Key, "session", t)
	stringAssertEqual(cookies[1].Value, "yes=no=yes=3;or=4", t)

}