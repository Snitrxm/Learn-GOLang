package productController 

import (
	"projeto/models"
	"github.com/gofiber/fiber/v2"
	"projeto/database"
	"errors"
)

type Product struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Price float32 `json:"price"`
}

func FindProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)

	if product.ID == 0 {
		return errors.New("Product not found")
	}

	return nil
}

func CreateProductResponse(productModel models.Product) Product {
	return Product{
		ID: productModel.ID,
		Name: productModel.Name,
		Price: productModel.Price,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	database.Database.Db.Create(&product);

	productResponse := CreateProductResponse(product)

	return c.Status(201).JSON(productResponse);
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id");

	if err != nil {
		return c.Status(400).SendString("Params needs to be a integer")
	}

	var product models.Product

	if err := FindProduct(id, &product); err != nil {
		return c.Status(404).SendString(err.Error())
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error());
	}

	return c.Status(200).SendString("Produt deleted")


}