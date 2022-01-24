package data

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

const OK = 0
const ID_TAKEN = 2

type Entry struct {
	gorm.Model
	EntryId  uint
	AuthorId string
	Text     string
}

type User struct {
	gorm.Model
	UserId     string
	Password   string
	Name       string
	Email      string
	ShowName   bool
	JoinedDate string
}

type Comment struct {
	gorm.Model
	EntryId  uint
	AuthorId string
	Text     string
	Likes    int
	Dislikes int
}

var db *gorm.DB

func InitDB() {
	ClearDatabase()
	var err error
	db, err = gorm.Open(sqlite.Open("web.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Entry{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Comment{})

	// Create
	//db.Create(&Entry{Code: "D42", Price: 100})

	// Read
	//var product Entry
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	////db.Model(&product).Updates(Entry{Price: 200, Code: "F42"}) // non-zero fields
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

func ValidateUser(userId string) bool {
	var user User
	res := db.Where(User{UserId: userId}).First(&user)
	if res.RowsAffected != 0 {
		return false
	}
	return true
}

func RegisterUser(userId, password, email string) {
	res := db.Create(&User{
		UserId:     userId,
		Name:       "",
		Email:      email,
		Password:   password,
		ShowName:   false,
		JoinedDate: time.Now().String()})

	if res.Error != nil {
		fmt.Printf("Error. Cannot create user %s: %s", userId, res.Error)
		return
	}

	fmt.Printf("User %s created", userId)
}

func AuthenticateUser(userId, password string) bool {
	if !ValidateUser(userId) {
		return false
	}

	var user User
	res := db.Where(User{UserId: userId, Password: password}).First(&user)
	if res.RowsAffected != 0 {
		return false
	}
	return true
}

func ModifyUser(name, password, email string, showName bool) {
	var user *User
	db.First(&user)

	if user == nil {
		return
	}

	user.Name = name
	user.Email = email
	user.Password = password
	user.ShowName = showName
	db.Save(&user)
}

func GetEntry(id string) Entry {
	var entry Entry
	db.First(&entry, "id = ?", id) // find product with code D42
	return entry
}

func GetEntriesForAuthor(authorId string) []Entry {
	var entries []Entry
	db.Where("authorId LIKE ?", authorId).Find(&entries)
	return entries
}
