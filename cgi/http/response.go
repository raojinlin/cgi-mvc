package http

import (
	"bytes"
	"strconv"
)

const MAX_COOKIES int = 10

type Response struct {
	statusCode    int
	contentType   string
	contentLength int
	Headers       map[string]string
	Body          *bytes.Buffer
	cookies       []*Cookie
	cookieIndex   int
}

func (res *Response) Len() int {
	return res.contentLength
}

func (res *Response) SetLen(length int ) {
	res.contentLength = length
	res.SetHeader("Content-Length", strconv.Itoa(res.contentLength))
}

func (res *Response) SetContentType(contentType string) {
	res.contentType = contentType
	res.SetHeader("Content-Type", contentType)
}

func (res *Response) GetContentType() string  {
	return res.contentType
}

func (res *Response) SetHeader(k, v string) {
	res.Headers[k] = v
}

func (res *Response) GetHeader(k string) string  {
	return res.Headers[k]
}

func (res *Response) SetStatus(status int) {
	res.statusCode = status
	res.SetHeader("Status", strconv.Itoa(res.statusCode) + " " + Status[res.statusCode])
}

func (res *Response) GetStatus() int  {
	return res.statusCode
}

func (res *Response) SetCookie(cookie *Cookie)  {
	if len(res.cookies) <= res.cookieIndex {
		return
	}

	res.cookies[res.cookieIndex] = cookie
	res.cookieIndex++
}

func (res *Response) GetCookies() []*Cookie {
	return res.cookies
}

func (res *Response) Redirect(location string) {
	res.SetHeader("Location", location)
}

func NewResponse(statusCode int) *Response {
	return &Response{
		statusCode:    statusCode,
		contentType:   "",
		contentLength: 0,
		Headers:       make(map[string]string),
		Body:          nil,
		cookies:       make([]*Cookie, MAX_COOKIES),
		cookieIndex:   0,
	}
}
