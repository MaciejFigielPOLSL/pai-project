package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func InitDB() {
	//ClearDatabase()
	var err error
	db, err = gorm.Open(sqlite.Open("web.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Comment{})

	// Create
	//db.Create(&Article{Code: "D42", Price: 100})

	// Read
	//var product Article
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	////db.Model(&product).Updates(Article{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)
}

func ClearDatabase() {
	e := os.Remove("web.db")
	if e != nil {
		log.Fatal(e)
	}
}
