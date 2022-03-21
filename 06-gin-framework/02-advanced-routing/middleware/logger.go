package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Logger() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		log.Printf("[CRUDAPP] %s %s (%d)", ctx.Request.Method, ctx.Request.URL.String(), ctx.GetInt("RequestId"))
	}
}
