package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"projeto/database"
	"projeto/modules/user/routes"
	"projeto/modules/product/routes"
	"projeto/modules/order/routes"
)


func main(){
	database.ConnectDb();
	app := fiber.New();

	userRoutes.SetupUserRoutes(app)
	productRoutes.SetupProductRoutes(app)
	orderRoutes.SetupOrderRoutes(app)

	log.Fatal(app.Listen(":8080"))
}