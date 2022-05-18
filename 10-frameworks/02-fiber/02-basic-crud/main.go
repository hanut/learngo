package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtMiddleware "github.com/gofiber/jwt/v3"
	"github.com/gofiber/template/html"
	"hanutsingh.in/learngo/frameworks/fiber/02/controllers"
)

var validate = validator.New()

func main() {
	// Create a new engine
	tplEngine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: tplEngine,
	})

	app.Use(logger.New())

	app.Use(func(c *fiber.Ctx) error {
		sess, err := SessionStore.Get(c)
		if err != nil {
			fmt.Println("SESSION WARNING:", err.Error())
		} else {
			fmt.Println("New Session ? ", sess.Fresh())
		}
		return c.Next()
	})

	// JWT Middleware
	app.Use("/users", jwtMiddleware.New(jwtMiddleware.Config{
		SigningKey: []byte("replace this with an actual key"),
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/webapp", 301)
	})

	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("./assets"),
	}))

	app.Route("/webapp", controllers.WebappController)
	app.Route("/users", controllers.UserController)
	app.Route("/auth", controllers.AuthenticationController)

	fmt.Println("API Server listening on :3000")
	app.Listen(":3000")
}
