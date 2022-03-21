package routes

import (
	"hanut/learngo/gin/crudapp/libs/auth"

	"github.com/gin-gonic/gin"
)

var RoleAdmin = "Admin"

var ManageAccount = auth.WithJWT(func(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}, &RoleAdmin)

var ManageMyProfile = auth.WithJWT(func(ctx *gin.Context) {
	ctx.JSON(200, "OK")
}, nil)
