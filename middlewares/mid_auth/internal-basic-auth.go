package mid_auth

import "github.com/gin-gonic/gin"

func InternalBasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"inter": "com",
	})
}
