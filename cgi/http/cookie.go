package http

import (
	"fmt"
	"time"
)

type Cookie struct {
	Key string
	Value string
	Domain string
	Path string
	Expires *time.Time
	MaxAge int
	HttpOnly bool
	SameSite bool
	Secure bool
}

func (c *Cookie) SetDomain(domain string)  {
	c.Domain = domain
}

func (c *Cookie) SetExpiresTime(expires *time.Time)  {
	c.Expires = expires
}

func (c *Cookie) String() string {
	cookieString := fmt.Sprintf("%s=%s; Domain=%s; Path=%s;",
		c.Key, c.Value, c.Domain, c.Path)

	if c.Expires != nil {
		cookieString += " " + c.Expires.Format(time.UnixDate) + ";"
	}

	if c.HttpOnly {
		cookieString += " HttpOnly;"
	}

	if c.Secure {
		cookieString += " Secure;"
	}

	if c.MaxAge != 0 {
		cookieString += " MaxAge=" + string(c.MaxAge) + ";"
	}

	return cookieString
}

func NewCookie(domain string) *Cookie {
	return &Cookie{
		Domain: domain,
	}
}
