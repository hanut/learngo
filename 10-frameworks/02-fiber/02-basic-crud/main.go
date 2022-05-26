package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"hanutsingh.in/learngo/frameworks/fiber/02/controllers"
)

const ConnTimeout = time.Second * 30

var validate = validator.New()

func main() {
	// Create a new engine
	tplEngine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:             tplEngine,
		AppName:           "Simple User Service",
		Concurrency:       100, // Limits this service to a maximum of 100 concurrent connections
		ReadTimeout:       ConnTimeout,
		WriteTimeout:      ConnTimeout,
		IdleTimeout:       ConnTimeout,
		StreamRequestBody: true,
	})

	app.Use(logger.New())

	// Initialize the Session
	InitSession(app)

	// Initialize the JWT Middleware
	InitJwt(app)

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
