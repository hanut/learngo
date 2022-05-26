package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
)

// Create a new session store
var SessionStore *session.Store

func InitSession(app *fiber.App) {
	fmt.Println("Setting up session store and middleware...")
	SessionStore = session.New(session.Config{
		Storage: redis.New(redis.Config{
			Host: "localhost",
			Port: 6379,
		}),
	})

	app.Use("/webapp/dashboard", validSessionMiddleware)
	app.Use("/api/dashboard", validSessionMiddleware)
}

func validSessionMiddleware(c *fiber.Ctx) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		fmt.Println("SESSION WARNING:", err.Error())
	} else {
		fmt.Println("New Session ? ", sess.Fresh())
	}
	return c.Next()
}
