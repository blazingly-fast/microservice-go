package data

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float32 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

type ProductDB struct {
	conn *pgx.Conn
}

func NewProductDB(conn *pgx.Conn) *ProductDB {
	return &ProductDB{conn}
}

// Products defines a slice of Product
type Products []*Product

// GetProducts returns all products from the database
func (db *ProductDB) GetProducts() []*Product {
	var products_list []*Product
	rows, err := db.conn.Query(context.Background(), "select * from products")
	if err != nil {
		log.Fatalf("failed to get products: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.SKU)
		if err != nil {
			log.Fatalf("failed to list products: %v", err)
		}
		products_list = append(products_list, &product)
	}
	return products_list
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func (db *ProductDB) GetProductByID(id int) (*Product, error) {
	sql := fmt.Sprintf("select * from products where id=%d", id)
	rows, err := db.conn.Query(context.Background(), sql)
	if err != nil {
		log.Fatalf("failed to get product: %v", err)
	}
	product := Product{}
	for rows.Next() {
		rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.SKU)
	}
	return &product, nil
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func (db *ProductDB) UpdateProduct(p Product) error {
	sql := fmt.Sprintf(
		"update products set name='%s', description='%s', price='%f', sku='%s' where id=%d", p.Name, p.Description, p.Price, p.SKU, p.ID)
	tx, err := db.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
	}
	_, err = tx.Exec(context.Background(), sql)
	if err != nil {
		log.Fatal("tx.Exec failed: %v", err)
	}
	tx.Commit(context.Background())
	return nil
}

// AddProduct adds a new product to the database
func (db *ProductDB) AddProduct(p Product) {
	sql := fmt.Sprintf("insert into products(name, description, price, sku) values('%s','%s','%f','%s')", p.Name, p.Description, p.Price, p.SKU)
	tx, err := db.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
	}
	_, err = tx.Exec(context.Background(), sql)
	if err != nil {
		log.Fatal("tx.Exec failed: %v", err)
	}
	tx.Commit(context.Background())
}

// DeleteProduct deletes a product from the database
func (db *ProductDB) DeleteProduct(id int) error {
	sql := fmt.Sprintf("delete from products where id=%d", id)
	tx, err := db.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
	}
	_, err = tx.Exec(context.Background(), sql)
	if err != nil {
		log.Fatal("tx.Exec failed: %v", err)
	}
	tx.Commit(context.Background())
	return nil
}
