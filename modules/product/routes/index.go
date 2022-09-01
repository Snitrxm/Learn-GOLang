package productRoutes

import (
	"github.com/gofiber/fiber/v2"
	"projeto/modules/product/controller"
)

func SetupProductRoutes(app *fiber.App){
	app.Post("/product", productController.CreateProduct)
	app.Delete("/product/:id", productController.DeleteProduct)
}