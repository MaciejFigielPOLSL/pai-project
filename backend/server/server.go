package server

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sources/data"
)

const (
	userkey = "user"
)

type fullEntry struct {
	Entry    data.Article   `json:"entry"`
	Comments []data.Comment `json:"comments"`
}

func Run() {
	r := engine()
	r.Use(gin.Logger())
	if err := engine().Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func engine() *gin.Engine {
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	r.POST("/api/register", registerUser)
	r.POST("/api/login", login)
	r.GET("/api/logout", logout)
	r.GET("/api/modify", authRequired, modifyUser)
	r.POST("/api/clear", authRequired, clearAllArticles)

	r.GET("/articles/page/:page", getArticlesForPage)
	r.GET("/articles/:entryId", getFullArticleData)
	r.GET("/articles/", getAllArticles)
	r.POST("/articles/", authRequired, addArticle)

	r.GET("/comment/get/:entryId", authRequired, getArticleComments)
	r.GET("/comment/like/:commentId", addLike)
	r.GET("/comment/dislike/:commentId", addLike)
	r.POST("/comment", authRequired, addComment)
	r.GET("/comment", getComments)

	private := r.Group("/api/private")
	private.Use(authRequired)
	{
		private.GET("/me", me)
		private.GET("/status", status)
		private.POST("/deleteme", deleteUser)
	}
	return r
}

// authRequired is a simple middleware to check the session
func authRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		fmt.Printf("User not logged in\n")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func clearAllArticles(c *gin.Context) {
	data.DeleteComments()
	data.DeleteArticles()
}
