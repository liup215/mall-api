package auth

import (
	"time"

	"github.com/appleboy/gin-jwt"
)

func New(c *Config) *jwt.GinJWTMiddleware {
	if c == nil {
		c = &Config{}
	}

	if c.Realm == "" {
		c.Realm = "mall"
	}

	if c.Key == "" {
		c.Key = "mallmallmall"
	}
	if c.TokenLookup == "" {
		c.TokenLookup = "header: Authorization, query: token, cookie: jwt"
	}

	if c.TokenHeadName == "" {
		c.TokenHeadName = "Bearer"
	}

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:   c.Realm,
		Key:     []byte(c.Key),
		Timeout: 24 * 30 * time.Hour,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		// TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenLookup: c.TokenLookup,
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		// TokenHeadName: "Bearer",
		TokenHeadName: c.TokenHeadName,

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	}

	return authMiddleware
}
