package http

type Request struct {
	Method string
	Uri string
	Headers map[string]string
	ContentType string
	ContentLength int
	RemoteAddr string
	RemotePort int
	QueryString string
	Scheme string
	QueryParams map[string]string
	Body string
	PostPrams map[string]string
	Cookies []*Cookie
	Params map[string]string
}

func (req *Request) GetHeader(field string) string  {
	return req.Headers[field]
}

func (req *Request) GetCookie(key string) *Cookie {
	for _, cookie := range req.Cookies {
		if cookie != nil && cookie.Key == key {
			return cookie
		}
	}

	return nil
}

func (req *Request) SetBody(body string)  {
	req.Body = body
	req.PostPrams = parseQueryParamsFromString(body)
}

func (req *Request) SetParams(params map[string]string)  {
	req.Params = params
}


func NewRequest(params map[string]string, scheme string) *Request {
	headers := getHeadersFromEnvMap(params, scheme)
	return &Request{
		Method:        params["REQUEST_METHOD"],
		Uri:           params["REQUEST_URI"],
		RemoteAddr:    params["REMOTE_ADDR"],
		RemotePort:    parseInt(params["REMOTE_PORT"]),
		Headers:       headers,
		ContentType:   params["CONTENT_TYPE"],
		ContentLength: parseInt(params["CONTENT_LENGTH"]),
		Scheme:        scheme,
		QueryString:   params["QUERY_STRING"],
		QueryParams:   parseQueryParamsFromString(params["QUERY_STRING"]),
		Cookies:       parseCookies(headers["Cookie"]),
	}
}