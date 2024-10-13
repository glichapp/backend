package middleware

import "github.com/gin-gonic/gin"

// CORS ...
func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement me
		ctx.Next()
	}
}
