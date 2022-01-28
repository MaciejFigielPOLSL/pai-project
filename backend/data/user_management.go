package data

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username   string
	Password   string
	Name       string
	Email      string
	ShowName   bool
	JoinedDate string
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

func DeleteUser(userId string) {
	db.Delete(&User{}, userId)
}
