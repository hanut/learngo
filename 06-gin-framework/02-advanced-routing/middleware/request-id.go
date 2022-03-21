package middleware

import (
	"github.com/gin-gonic/gin"
)

// Middlewares can be used to modify incoming requests,
// add data to the requests or generally perform some tasks
// on the basis of the request

// RequestIdMiddleware returns a new handler function that adds
// a sequential generated request id to the incoming requests
func RequestIdMiddleware() func(ctx *gin.Context) {
	rid := 0
	return func(ctx *gin.Context) {
		// Increment the counter
		rid++
		ctx.Set("RequestId", rid)
	}
}
