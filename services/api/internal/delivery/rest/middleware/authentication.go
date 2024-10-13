package middleware

import "github.com/gin-gonic/gin"

// Authentication ...
func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement me
		ctx.Next()
	}
}
