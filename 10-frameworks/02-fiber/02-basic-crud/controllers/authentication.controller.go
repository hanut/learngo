package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"hanutsingh.in/learngo/frameworks/fiber/02/database"
	utils "hanutsingh.in/learngo/frameworks/fiber/02/utils"
)

type LoginForm struct {
	UserId   string `json:"uid" validate:"required"`
	Password string `json:"pass" validate:"required,min=8,max=16,ascii"`
}

func AuthenticationController(r fiber.Router) {

	// Route handler for the main index page before login
	r.Post("/form", func(c *fiber.Ctx) error {
		var lf LoginForm
		if err := c.BodyParser(&lf); err != nil {
			fmt.Println("Error:", err)
			return fiber.NewError(400, err.Error())
		}
		if err := validate.Struct(lf); err != nil {
			return fiber.NewError(400, err.Error())
		}
		// Check if the userid+pass combo is correct
		u, err := database.UserStore.GetUser(lf.UserId)
		if err != nil {
			return fiber.NewError(401, "Invalid username or password")
		}
		if u.Password != lf.Password {
			return fiber.NewError(401, "Invalid username or password")
		}

		expat := time.Now().Add(time.Hour * 72).Unix()
		// Create the Claims
		claims := jwt.MapClaims{
			"name": fmt.Sprintf("%s %s", u.FirstName, u.LastName),
			"role": u.Role,
			"exp":  expat,
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		sess, err := utils.SessionStore.Get(c)
		sess.Set("username", lf.UserId)
		sess.SetExpiry(time.Second * 60)
		return c.JSON(fiber.Map{"token": t, "expiry": expat})
	})

}
