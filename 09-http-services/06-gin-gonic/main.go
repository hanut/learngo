package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(404, gin.H{
			"name":   "aniruddh",
			"age":    38,
			"salary": 500000,
		})
	})
	r.GET("/", func(c *gin.Context) {
		fmt.Println("Params", c.Params)
		fmt.Println("Name", c.Query("name"))
		fmt.Println("Headers", c.Request.Header)

		c.Writer.Write([]byte("<h1>Hi</h1>"))
	})
	r.Run("localhost:8085")
}
