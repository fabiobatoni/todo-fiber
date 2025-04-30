package users

import (
	"net/http"

	"github.com/fabiobatoni/todo-fiber/db"
	"github.com/gofiber/fiber/v2"
)

func getAll(c *fiber.Ctx) error {
	var documents []User

	err := db.Find("users", &documents)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(documents)
}
