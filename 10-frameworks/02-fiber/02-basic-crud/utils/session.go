package utils

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
			Host:     "localhost",
			Port:     6379,
			Database: 3,
		}),
	})

	app.Use("/webapp/dashboard", validSessionMiddleware)
}

func validSessionMiddleware(c *fiber.Ctx) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		return fiber.ErrUnauthorized
	}
	if sess.Fresh() {
		return c.Redirect("/webapp/")
	}
	fmt.Println("Session Id:", sess.ID())
	return c.Next()
}
