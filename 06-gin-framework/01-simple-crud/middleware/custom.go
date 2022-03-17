package middleware

import (
	"github.com/gin-gonic/gin"
)

var CustomMiddleware gin.HandlerFunc = func(ctx *gin.Context) {
	// This is a custom handler implementation that adds custom data to the
	// ctx value

	ctx.Request.
}
