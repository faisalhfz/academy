package routes

import (
	"user-management/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router, userHandler *handler.UserHandler) {
	r := app.Group("api/v1")

	r.Post("/user", userHandler.CreateUser)
	r.Get("/users", userHandler.GetList)
	r.Get("/user/:id", userHandler.GetUser)
	r.Put("/user/:id", userHandler.UpdateUser)
	r.Delete("/user/:id", userHandler.DeleteUser)
}
