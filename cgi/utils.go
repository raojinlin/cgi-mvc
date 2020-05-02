package cgi

import (
	"fmt"
	"io"
	"iogo/cgi/http"
	"os"
	"strings"
)

func GetEnvironMapFromEnvirons(environs []string) map[string]string {
	env := make(map[string]string)

	for i := 0; i < len(environs); i++ {
		envKeyValueArray := strings.Split(environs[i], "=")
		if len(envKeyValueArray) == 2 {
			env[envKeyValueArray[0]] = envKeyValueArray[1]
		} else if len(envKeyValueArray) > 2 {
			env[envKeyValueArray[0]] = strings.Join(envKeyValueArray[1:], "=")
		}
	}

	return env
}

func SetPostParamsToContext(ctx *http.Context) {
	if ctx.Request.Method != "POST" {
		return
	}

	var reader io.Reader
	var err error
	reader, err = os.OpenFile("/dev/stdin", os.O_RDONLY, 0)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bytes := make([]byte, ctx.Request.ContentLength)
	_, err = reader.Read(bytes)

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx.Request.SetBody(string(bytes))
}