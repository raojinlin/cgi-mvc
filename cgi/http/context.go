package http

import "iogo/cgi/config"

type Context struct {
	Request *Request
	Response *Response

	DocumentRoot string
	Prefix string

	ServerProtocol string
	ServerAddr string
	ServerPort int
	ServerSoftware string
	ServerSignature string

	GatewayInterface string
	Scheme string

	Config *config.Config
}

func NewContext(envs map[string]string, config *config.Config) *Context  {
	return &Context{
		DocumentRoot:     getEnvValue(envs, "CONTEXT_DOCUMENT_ROOT"),
		Prefix:           getEnvValue(envs, "CONTEXT_PREFIX"),
		GatewayInterface: getEnvValue(envs, "GATEWAY_INTERFACE"),
		ServerProtocol:   getEnvValue(envs, "SERVER_PROTOCOL"),
		ServerSignature:  getEnvValue(envs, "SERVER_SIGNATURE"),
		Scheme:           getEnvValue(envs, "REQUEST_SCHEME"),
		ServerSoftware:   getEnvValue(envs, "SERVER_SOFTWARE"),
		Request:          NewRequest(envs, envs["REQUEST_SCHEME"]),
		Response:         NewResponse(200),
		Config:           config,
	}
}
