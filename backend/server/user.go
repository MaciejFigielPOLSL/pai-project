package server

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sources/data"
	"strings"
)

// @Summary Register new user
// @Description Add new user to system
// @Tags user
// @Param        username  formData      string  true  "User login"
// @Param        email  formData      string  true  "Email address"
// @Param        password  formData      string  true  "User password"
// @Success 200 {object} map[string]interface{}
// @Failure 406 {object} map[string]interface{}
// @Router /api/register [post]
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

// @Summary Login
// @Description Allows user to log into the system
// @Tags user
// @Param        username  formData      string  true  "User login"
// @Param        password  formData      string  true  "User password"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/login [post]
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

// @Summary Logout
// @Description Allows user to log out from the system
// @Tags user
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/logout [post]
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

// @Summary Current user
// @Description Returns informations about current user
// @Tags user
// @Param        username  formData      string  true  "User login"
// @Param        password  formData      string  true  "User password"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/private/me [post]
func me(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get(userkey)
	user := data.GetUser(username.(string))
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// @Summary Current user
// @Description Returns informations about current user
// @Tags user
// @Param        username  formData      string  true  "User login"
// @Param        password  formData      string  true  "User password"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /api/private/me [post]
func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

// @Summary Checks if user is logged
// @Description Returns info if user is logged in
// @Tags user
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/private/me [post]
func isLoggedIn(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.JSON(http.StatusOK, gin.H{"logged": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"logged": true})
}

// @Summary Modify user
// @Description Allows user to modify its data
// @Tags user
// @Param        username  formData      string  true  "User login"
// @Param        name  formData      string  true  "User name"
// @Param        password  formData      string  true  "User password"
// @Param        email  formData      string  true  "User email"
// @Param        showName  formData      bool  true  "Show name or not"
// @Produce json
// @Success 200 {string} string
// @Failure 401 {string} string
// @Failure 500 {string} string
// @Router /api/private/me [post]
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

// @Summary Delete user
// @Description Removes user from system
// @Tags user
// @Router /api/private/deleteme [post]
func deleteUser(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get(userkey).(string)

	fmt.Printf("Deleting user %s data\n", username)
	data.DeleteUser(username)
}
