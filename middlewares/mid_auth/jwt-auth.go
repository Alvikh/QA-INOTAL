package mid_auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		autH := strings.TrimSpace(authHeader)
		//log.Println("auth :", autH)
		//log.Println("len(auth) :", len(autH))

		if len(autH) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if len(autH) == 6 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := NewJWTService().ValidateToken(tokenString)
		if err == nil {
			if token.Valid {
				//claims := token.Claims.(jwt.MapClaims)
				/*
					log.Println("Claims[Name]: ", claims["name"])
					log.Println("Claims[Admin]: ", claims["admin"])
					log.Println("Claims[Issuer]: ", claims["iss"])
					log.Println("Claims[IssuedAt]: ", claims["iat"])
					log.Println("Claims[ExpiresAt]: ", claims["exp"])
				*/
			} else {
				//log.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			//log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
