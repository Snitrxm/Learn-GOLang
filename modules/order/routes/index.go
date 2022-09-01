package orderRoutes

import (
	"github.com/gofiber/fiber/v2"
	"projeto/modules/order/controller"
)

func SetupOrderRoutes(app *fiber.App){
	app.Post("/order", orderController.CreateOrder)
	app.Get("/order", orderController.GetOrders)
}