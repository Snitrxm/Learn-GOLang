package orderController

import (
	"github.com/gofiber/fiber/v2"
	"projeto/database"
	"projeto/models"
	"projeto/modules/user/controller"
	"projeto/modules/product/controller"
)

type Order struct {
	ID uint `json:"id"`
	Product productController.Product `json:"product"`
	User userController.User `json:"user"`
}

func CreateOrderResponse(orderModel models.Order, user userController.User, product productController.Product) Order {
	return Order{
		ID:         orderModel.ID,
		Product:   	product,
		User:      	user,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	var user models.User
	if err := userController.FindUser(order.UserRefer, &user); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	var product models.Product
	if err := productController.FindProduct(order.ProductRefer, &product); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	database.Database.Db.Create(&order);

	userResponse := userController.CreateResponseUser(user);
	productResponse := productController.CreateProductResponse(product);

	orderResponse := CreateOrderResponse(order, userResponse, productResponse)

	return c.Status(201).JSON(orderResponse);
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}

	database.Database.Db.Find(&orders)

	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product

		database.Database.Db.Find(&user, "id = ?", order.UserRefer);
		database.Database.Db.Find(&product, "id = ?", order.ProductRefer);
		responseOrder := CreateOrderResponse(order, userController.CreateResponseUser(user), productController.CreateProductResponse(product))
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.JSON(responseOrders)
}