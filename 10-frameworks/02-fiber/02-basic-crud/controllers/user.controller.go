package controllers

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"hanutsingh.in/learngo/frameworks/fiber/02/database"
	models "hanutsingh.in/learngo/frameworks/fiber/02/database/models"
)

var (
	// UserStore is the mock datastore for user documents
	UserStore = database.UserStore

	// validate is the validator instance for all routes in the user
	// controller
	validate = validator.New()
)

// UserController is a router function that sets up the CRUD
// endpoints for the "/users" route
func UserController(r fiber.Router) {

	// List all users in the datastore
	r.Get("/", func(c *fiber.Ctx) error {
		if len(UserStore) == 0 {
			return fiber.NewError(404, "No users in the UserStore as yet")
		}
		return c.JSON(UserStore)
	})

	// List a specific user by id from the datastore
	r.Get("/:id", func(c *fiber.Ctx) error {
		if len(UserStore) == 0 {
			return fiber.NewError(404, "No users in the UserStore as yet")
		}
		uid := c.Params("id", "")
		if uid == "" {
			return fiber.ErrBadRequest
		}
		u, err := UserStore.GetUser(uid)
		if err != nil {
			return fiber.NewError(400, err.Error())
		}
		return c.Status(200).JSON(u)
	})

	// Add a new user to the datastore if it doesnt already exist
	r.Post("/", func(c *fiber.Ctx) error {
		var user models.User
		if err := c.BodyParser(&user); err != nil {
			return fiber.NewError(400, err.Error())
		}
		errors := func() string {
			var errors string
			err := validate.Struct(user)
			if err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					errors += fmt.Sprintf("%s\t%s\t%s\n", err.StructNamespace(), err.Tag(), err.Param())
				}
			}
			return errors
		}()
		if errors != "" {
			return fiber.NewError(400, errors)
		}
		uid, err := UserStore.AddUser(user)
		if err != nil {
			return fiber.NewError(400, err.Error())
		}
		return c.JSON(fiber.Map{uid: uid})
	})

	// Remove a user from the datastore
	r.Delete("/:id", func(c *fiber.Ctx) error {
		if len(UserStore) == 0 {
			return fiber.NewError(404, "No users in the UserStore as yet")
		}
		uid := c.Params("id", "")
		if uid == "" {
			return fiber.ErrBadRequest
		}
		err := UserStore.RemoveUser(uid)
		if err != nil {
			return fiber.NewError(400, err.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

}
