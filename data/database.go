package data

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type Entry struct {
	gorm.Model
	EntryId  uint
	AuthorId string
	Title    string
	Text     string
	AddDate  time.Time
}

type Comment struct {
	gorm.Model
	EntryId  uint
	AuthorId string
	Text     string
	Likes    int32
	Dislikes int32
}

type User struct {
	gorm.Model
	Username   string
	Password   string
	Name       string
	Email      string
	ShowName   bool
	JoinedDate string
}

var db *gorm.DB

func InitDB() {
	//ClearDatabase()
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

func GetUser(userId string) User {
	var user User
	if !UserExists(userId) {
		return user
	}
	db.Where(User{Username: userId}).First(&user)
	return user
}

func UserExists(userId string) bool {
	var user User
	res := db.Where(User{Username: userId}).First(&user)
	return !errors.Is(res.Error, gorm.ErrRecordNotFound)
}

func RegisterUser(userId, password, email string) {
	res := db.Create(&User{
		Username:   userId,
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
	if UserExists(userId) {
		return false
	}

	var user User
	res := db.Where(User{Username: userId, Password: password}).First(&user)
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

var entriesPerPage = 10

func GetAllEntries() []Entry {
	var entries []Entry
	db.Order("add_date").Find(&entries)
	return entries
}

func GetEntriesOnPage(page int) []Entry {
	var entries []Entry
	result := db.Order("add_date").Offset(page * entriesPerPage).Limit(entriesPerPage).Find(&entries)

	if result.Error != nil {

	}
	return entries
}

func AddEntry(username string, title, text string) {
	entry := Entry{
		AuthorId: username,
		Title:    title,
		Text:     text,
		AddDate:  time.Now(),
	}
	db.Create(&entry)
}

func AddComment(entryId uint, authorId, text string) {
	comment := Comment{
		EntryId:  entryId,
		AuthorId: authorId,
		Text:     text,
		Likes:    0,
		Dislikes: 0,
	}
	db.Create(&comment)
}

func AddLike(commentId int32) {
	var comment Comment
	db.First(&comment, commentId)
	comment.Likes++
	db.Save(&comment)
}

func AddDislike(commentId int32) {
	var comment Comment
	db.First(&comment, commentId)
	comment.Dislikes++
	db.Save(&comment)
}

func GetFullEntry(id string) (Entry, []Comment) {
	return GetEntry(id), GetEntryComments(id)
}

func GetEntry(id string) Entry {
	var entry Entry
	db.First(&entry, "id = ?", id)
	return entry
}

func GetEntryComments(entryId string) []Comment {
	var comments []Comment
	db.Find(&comments, "entry_id = ?", entryId)
	return comments
}

func GetEntriesForAuthor(authorId string) []Entry {
	var entries []Entry
	db.Where("author_id LIKE ?", authorId).Find(&entries)
	return entries
}

func GetAllComments() []Comment {
	var comments []Comment
	db.Find(&comments)
	return comments
}
