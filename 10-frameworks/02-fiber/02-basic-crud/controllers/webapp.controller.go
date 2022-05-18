package controllers

import "github.com/gofiber/fiber/v2"

func WebappController(r fiber.Router) {

	// Route handler for the main index page before login
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).Render("pages/home", fiber.Map{
			"Title": "Admin Portal",
		}, "layouts/public")
	})

	// Route handler for the main index page before login
	r.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.Status(200).Render("pages/dashboard", fiber.Map{
			"Title": "Admin Portal | Dashboard",
		}, "layouts/dashboard")
	})

}
