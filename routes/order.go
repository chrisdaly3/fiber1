package routes

import (
	"errors"
	"time"

	"example.com/fiber1/database"
	"example.com/fiber1/models"
	"github.com/gofiber/fiber/v2"
)

/* Design for orders DB response
{
	id: 1,
	user: {
		id: 23,
		first_name: "John",
		last_name: "Doe",
	},
	product: {
		id: 3,
		name: "item_ordered",
		serial_number: "xyz23913",
	}
}
*/

type OrderSerializer struct {
	ID        uint              `json:"id"`
	User      UserSerializer    `json:"user"`
	Product   ProductSerializer `json:"product"`
	CreatedAt time.Time         `json:"order_date"`
}

func CreateOrderResponse(order models.Order, user UserSerializer, product ProductSerializer) OrderSerializer {
	return OrderSerializer{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	userResponse := CreateResponseUser(user)
	productResponse := CreateProductResponse(product)
	orderResponse := CreateOrderResponse(order, userResponse, productResponse)

	return c.Status(200).JSON(orderResponse)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	ordersResponse := []OrderSerializer{}

	for _, order := range orders {
		var user models.User
		var product models.Product
		database.Database.Db.Find(&user, "id = ?", order.UserRefer)
		database.Database.Db.Find(&product, "id = ?", order.ProductRefer)
		orderResponse := CreateOrderResponse(order, CreateResponseUser(user), CreateProductResponse(product))
		ordersResponse = append(ordersResponse, orderResponse)
	}

	return c.Status(200).JSON(ordersResponse)
}

func FindOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please verify and ensure :id is an integer")
	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product

	database.Database.Db.First(&user, order.UserRefer)
	database.Database.Db.First(&product, order.ProductRefer)
	orderResponse := CreateOrderResponse(order, CreateResponseUser(user), CreateProductResponse(product))
	return c.Status(200).JSON(orderResponse)
}
