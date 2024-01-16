package main

import (
	// "database/sql"

	"database/sql"
	"fmt"
	"log"

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

type Product struct {
	ID    int
	Name  string
	Price int
}

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

	// Check the connection
	err = db.Ping()
	fmt.Printf("Ping's: %v\n", err)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected DB!")

	err = createProduct(&Product{Name: "GoDB test", Price: 665})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created Product Successful")
}

func createProduct(product *Product) error {

	_, err := db.Exec(
		"INSERT INTO public.products(name, price)VALUES ($1, $2);",
		product.Name,
		product.Price,
	)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
