package middleware

import "github.com/gin-gonic/gin"

// Recover ...
func Recover() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement me
		ctx.Next()
	}
}
