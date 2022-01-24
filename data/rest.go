package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
)

const (
	userkey = "user"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type user struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type fullEntry struct {
	Entry    Entry     `json:"entry"`
	Comments []Comment `json:"comments"`
}

func StartServer() {
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
	//router := gin.Default()
	//router.GET("/", getIndex)
	//router.POST("/register", registerUser)
	//
	//println("Running server on localhost:8080...")
	//router.Run("localhost:8080")
}

func engine() *gin.Engine {
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))
	r.POST("/register", registerUser)
	r.POST("/login", login)
	r.GET("/logout", logout)
	r.GET("/modify", AuthRequired, modifyUser)

	r.GET("/entries/:page", getEntries)
	r.GET("/entries/", getAllEntries)

	r.POST("/entry/", AuthRequired, addEntry)
	r.GET("/entry/:entryId", getFullEntry)

	r.GET("/comment/:entryId", AuthRequired, getEntryComments)
	r.POST("/comment", AuthRequired, addComment)
	r.GET("/comments", getComments)

	private := r.Group("/private")
	private.Use(AuthRequired)
	{
		private.GET("/me", me)
		private.GET("/status", status)
	}
	return r
}

var State = make(map[string]interface{})

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		fmt.Printf("User not logged in")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func registerUser(c *gin.Context) {
	//Get user name and password
	username := c.Request.FormValue("username")
	email := c.Request.FormValue("email")
	passwd := c.Request.FormValue("password")

	fmt.Printf("Trying to register user: %s, %s, %s", username, email, passwd)
	//Judge whether the user exists
	//Output state 1 present
	//No create user exists, save password and user name

	if UserExists(username) {
		//Registration status
		State["state"] = 1
		State["text"] = "This user already exists!"
		c.String(http.StatusNotAcceptable, "%v", State)
		return
	} else {
		//Add user if user does not exist
		RegisterUser(username, passwd, email)
		State["state"] = 0
		State["text"] = "Registration succeeded!"
		c.String(http.StatusOK, "%v", State)
		return
	}
}

// login is a handler that parses a form and checks for specific data
func login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if AuthenticateUser(username, password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
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
	user := GetUser(username.(string))
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

	if !UserExists(username) {
		fmt.Printf("Cannot find user to modify %s\n", username)
		c.String(http.StatusInternalServerError, "Failed to find user")
		return
	}

	ModifyUser(name, password, email, showName)
	c.String(http.StatusOK, "")
}

func getAllEntries(c *gin.Context) {
	entries := GetAllEntries()
	c.JSON(http.StatusOK, entries)
}

func getFullEntry(c *gin.Context) {
	//entryId, _ := strconv.ParseUint(c.Param("entryId"), 10, 32)
	entryId := c.Param("entryId")
	entry, comments := GetFullEntry(entryId)
	c.JSON(http.StatusOK, fullEntry{Entry: entry, Comments: comments})
}

func getEntries(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		fmt.Printf("Provided page '%s' is not a number\n", page)
		c.String(http.StatusInternalServerError, "Failed to parse page")
		return
	}
	if page <= 0 {
		fmt.Printf("Page '%s' must be positive number\n", page)
		c.String(http.StatusInternalServerError, "Failed to parse page")
		return
	}
	entries := GetEntriesOnPage(page - 1)
	c.JSON(http.StatusOK, entries)
}

type EntryJsonBody struct {
	Title string `form:"title" json:"title"`
	Text  string `form:"text" json:"text"`
}

func addEntry(c *gin.Context) {
	username := sessions.Default(c).Get(userkey).(string)
	//user := GetUser(username)
	var requestBody EntryJsonBody

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("%s\n", err)
	}

	fmt.Println(requestBody.Text)

	AddEntry(username, requestBody.Title, requestBody.Text)
}

func getEntryComments(c *gin.Context) {
	//entryId, _ := strconv.ParseUint(c.Param("entryId"), 10, 32)
	entryId := c.Param("entryId")
	c.JSON(http.StatusOK, GetEntryComments(entryId))
}

type CommentJsonBody struct {
	EntryId uint   `form:"entryId" json:"entryId"`
	Text    string `form:"text" json:"text"`
}

func addComment(c *gin.Context) {
	username := sessions.Default(c).Get(userkey).(string)
	//user := GetUser(username)
	var requestBody CommentJsonBody

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("%d: %s, %s\n", requestBody.EntryId, username, requestBody.Text)

	AddComment(requestBody.EntryId, username, requestBody.Text)
}

func getComments(c *gin.Context) {
	c.JSON(http.StatusOK, GetAllComments())
}
