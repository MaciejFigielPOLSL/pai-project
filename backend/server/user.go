package server

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sources/data"
	"strings"
)

func registerUser(c *gin.Context) {
	username := c.Request.FormValue("username")
	email := c.Request.FormValue("email")
	passwd := c.Request.FormValue("password")

	fmt.Printf("Trying to register user: %s, %s, %s\n", username, email, passwd)

	var RegisterState = make(map[string]interface{})

	if data.UserExists(username) {
		//Registration status
		RegisterState["state"] = 1
		RegisterState["text"] = "This user already exists!"
		c.String(http.StatusNotAcceptable, "%v", RegisterState)
		return
	} else {
		//Add user if user does not exist
		data.RegisterUser(username, passwd, email)
		RegisterState["state"] = 0
		RegisterState["text"] = "Registration succeeded!"
		c.String(http.StatusOK, "%v", RegisterState)
		return
	}
}

// login is a handler that parses a form and checks for specific data
func login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	if data.AuthenticateUser(username, password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	session.Set(userkey, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func me(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get(userkey)
	user := data.GetUser(username.(string))
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

func modifyUser(c *gin.Context) {
	username := c.PostForm("username")
	name := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	showName := c.PostForm("showName") == "true"

	if !data.UserExists(username) {
		fmt.Printf("Cannot find user to modify %s\n", username)
		c.String(http.StatusInternalServerError, "Failed to find user")
		return
	}

	fmt.Printf("Modifying user %s data\n", username)

	data.ModifyUser(name, password, email, showName)
	c.String(http.StatusOK, "")
}

func deleteUser(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get(userkey).(string)

	fmt.Printf("Deleting user %s data\n", username)
	data.DeleteUser(username)
}
