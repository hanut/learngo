package main

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		q := ctx.Query("q")
		ctx.JSON(200, q)
	})

	return r
}

func main() {
	r := SetupRouter()
	r.Run()
}
