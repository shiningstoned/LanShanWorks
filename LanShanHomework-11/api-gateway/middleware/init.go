package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(service interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Keys = make(map[string]interface{})
		ctx.Keys["userService"] = service
		ctx.Next()
	}
}
