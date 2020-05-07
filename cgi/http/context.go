package http

import (
	"iogo/cgi/config"
	"iogo/cgi/logger"
)

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
	Logger *logger.FileLogger
}

func (ctx *Context) Cookie(name string) *Cookie {
	return ctx.Request.GetCookie(name)
}

func (ctx *Context) Redirect(location string)  {
	ctx.Response.Redirect(location)
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
		Logger:           logger.NewFileLogger("Context", config.Logger.Output, logger.LogInfo),
	}
}
