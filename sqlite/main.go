// Quick Start from https://gorm.io/docs/
package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Ao usar o gorm.Model, o ID do registro é gerado automaticamente, por isso não é necessário passar o ID como parâmetro.
// O gorm.Model possui o campo deleted_at, ativando por padrão o uso do Soft Delete.
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("./db/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Print product
	fmt.Println("Id: ", product.ID, "Code: ", product.Code, "Price: ", product.Price)

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Print product
	fmt.Println("Id: ", product.ID, "Code: ", product.Code, "Price: ", product.Price)

	// Delete - delete product
	//db.Delete(&product, 1)
	db.Delete(&product, product.ID)
}
