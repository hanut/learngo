package controllers

import (
	"github.com/gofiber/fiber/v2"
	"hanutsingh.in/learngo/frameworks/fiber/02/utils"
)

func WebappController(r fiber.Router) {

	// Route handler for the main index page before login
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).Render("pages/home", fiber.Map{
			"Title":       "Admin Portal",
			"AuthDataKey": "authData",
		}, "layouts/public")
	})

	// Route handler for the main index page before login
	r.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.Status(200).Render("pages/dashboard", fiber.Map{
			"Title": "UMS Dashboard",
			"Page":  "List Users",
			"Users": UserStore,
		}, "layouts/dashboard")
	})

	// Route handler for the logging out of the webapp
	r.Get("/logout", func(c *fiber.Ctx) error {
		s, err := utils.SessionStore.Get(c)
		if err == nil {
			err = s.Destroy()
			if err != nil {
				return fiber.NewError(500, err.Error())
			}
		}
		return c.Redirect("/webapp/")
	})
}
