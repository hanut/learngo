package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appConf := fiber.Config{
		Concurrency: 16 * 1024,
		ReadTimeout: time.Second * 10,
	}
	app := fiber.New(appConf)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"message": "Hello World"})
	})

	fmt.Println("Listening on port 3000...")
	app.Listen(":3000")
}
