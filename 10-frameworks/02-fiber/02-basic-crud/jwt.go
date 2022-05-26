package main

import (
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v3"
)

func InitJwt(app *fiber.App) {
	jwtMW := jwtMiddleware.New(jwtMiddleware.Config{
		SigningKey: []byte("replace this with an actual key"),
	})
	// Apply the middleware to as many routes as needed
	app.Use("/users", jwtMW)
}
