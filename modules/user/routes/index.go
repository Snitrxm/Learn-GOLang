package userRoutes

import (
	"github.com/gofiber/fiber/v2"
	"projeto/modules/user/controller"
)


func SetupUserRoutes(app *fiber.App){
	app.Post("/user", userController.CreateUser)
	app.Get("/user", userController.GetUsers);
	app.Get("/user/:id", userController.GetUserById)
	app.Put("/user/:id", userController.UpdateUserById)
	app.Delete("/user/:id", userController.DeleteUserById)
}