package userController

import (
	"projeto/models"
	"github.com/gofiber/fiber/v2"
	"projeto/database"
	"errors"
)

type User struct {
	// Not the User model, just a serializer, like a interface in typescript
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID: userModel.ID,
		FirstName: userModel.FirstName,
		LastName: userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User;

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error());
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user);

	return c.Status(201).JSON(responseUser);
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users);
	responseUsers := []User{}

	for _, user := range users {
		responseUser := CreateResponseUser(user);
		responseUsers = append(responseUsers, responseUser);
	}

	return c.Status(200).JSON(responseUsers);
}

func FindUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return errors.New("User not found")
	}

	return nil
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id");

	if err != nil {
		return c.Status(400).JSON("Params should be a int");
	}

	var user models.User

	if err := FindUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error());
	}

	responseUser := CreateResponseUser(user);
	
	return c.Status(200).JSON(responseUser);
}

func UpdateUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id");

	if err != nil {
		return c.Status(400).JSON("Params should be a int");
	}

	var user models.User

	if err := FindUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error());
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser;

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error());
	}

	if updateData.FirstName != "" {
		user.FirstName = updateData.FirstName;
	}

	if updateData.LastName != "" {
		user.LastName = updateData.LastName;
	}

	database.Database.Db.Save(&user);

	responseUser := CreateResponseUser(user);

	return c.Status(200).JSON(responseUser);
}

func DeleteUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id");

	if err != nil {
		return c.Status(400).JSON("Params should be a int");
	}

	var user models.User

	if err := FindUser(id, &user); err != nil {
		return c.Status(404).JSON(err.Error());
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(500).JSON(err.Error());
	}

	return c.Status(200).SendString("User deleted");
}