package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"projeto/database"
	"projeto/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my API")
}

func setupUserRoutes(app *fiber.App){
	app.Get("/", welcome)
	app.Post("/user", routes.CreateUser)
	app.Get("/user", routes.GetUsers);
	app.Get("/user/:id", routes.GetUserById)
	app.Put("/user/:id", routes.UpdateUserById)
	app.Delete("/user/:id", routes.DeleteUserById)
}


func main(){
	database.ConnectDb();
	app := fiber.New();

	setupUserRoutes(app)

	log.Fatal(app.Listen(":8080"))
}