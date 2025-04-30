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

func getByID(c *fiber.Ctx) error {
	var document User

	err := db.FindByID("users", c.Params("id"), &document)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(document)
}
