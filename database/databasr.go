package database

import (
	"database/sql"
	"log"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func CreateProduct(product *Product) error {
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

func GetProduct(id int) (*Product, error) {
	var p Product
	row := db.QueryRow("SELECT id, name, price FROM products WHERE id = $1", id)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return &Product{}, err
	}
	return &p, nil
}

func GetAllProducts() ([]Product, error) {
	var products []Product
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	if err != nil {
		return nil, err
	}
	return products, nil
}

func UpdateProduct(id int, product *Product) (Product, error) {
	var p Product
	row := db.QueryRow(
		"UPDATE public.products SET name = $1, price = $2 WHERE id = $3 RETURNING id, name, price;",
		product.Name,
		product.Price,
	)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err

	}
	return p, nil
}

func DeleteProduct(id int) error {
	_, err := db.Exec("DELETE FROM public.products WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
