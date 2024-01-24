package main

import (
	// "database/sql"

	"database/sql"
	"fmt"
	"goConnectDB/database"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

var db *sql.DB

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection
	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	db = sdb
	defer db.Close()

	app := fiber.New()

	// Check the connection
	err = db.Ping()
	fmt.Printf("Ping's: %v\n", err)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the db variable in the database package
	database.SetDB(db)

	app.Get("/products", GetAllProductsHandler)
	app.Get("/product/:id", getProductHandler)
	app.Post("/product", createProductHandler)
	app.Put("/product/:id", updateProductHandler)
	app.Delete("/product/:id", deleteProductHandler)

	app.Listen(":8080")

	// // Call functions from the database package
	// products, err := database.GetAllProducts()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Retrieved Products:")
	// for _, p := range products {
	// 	fmt.Printf("ID: %d, Name: %s, Price: %d\n", p.ID, p.Name, p.Price)
	// }
}

func GetAllProductsHandler(c *fiber.Ctx) error {
	products, err := database.GetAllProducts()
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(products)
}

func getProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	product, err := database.GetProduct(productId)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(product)
}

func createProductHandler(c *fiber.Ctx) error {
	product := new(database.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	err := database.CreateProduct(product)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(product)
}

func updateProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	product := new(database.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	updatedProduct, err := database.UpdateProduct(productId, product)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(updatedProduct)
}

func deleteProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = database.DeleteProduct(productId)
	if err != nil {
		log.Fatal(err)
	}
	return c.SendStatus(204)
}
