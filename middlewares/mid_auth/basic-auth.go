package mid_auth

import "github.com/gin-gonic/gin"

// BasicAuth adalah middleware yang mengimplementasikan HTTP Basic Authentication.
// Fungsi ini mengembalikan middleware yang akan memvalidasi header Authorization di setiap request.
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"user-rsud": "password123", // Username dan password yang diotorisasi
	})
}
