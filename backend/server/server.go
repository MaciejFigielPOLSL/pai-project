package server

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sources/data"
)

const (
	userkey = "user"
)

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
	//r.Use(cors.New(cors.Config{
	//	AllowedOrigins:   []string{"http://*"},
	//	AllowedMethods:   []string{"GET", "POST"},
	//	AllowedHeaders:   []string{"Origin"},
	//	ExposedHeaders:   []string{"Content-Length"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))
	r.Use(CORSMiddleware())

	r.Use(static.Serve("/", static.LocalFile("./index/dist", true)))

	r.POST("/api/register", registerUser)
	r.POST("/api/login", login)
	r.POST("/api/logout", authRequired, logout)
	r.POST("/api/modify", authRequired, modifyUser)
	r.GET("/api/status", isLoggedIn)
	r.POST("/api/clear", authRequired, clearAllArticles)

	r.GET("/articles/page/:page", getArticlesForPage)
	r.GET("/articles/:articleId", getFullArticleData)
	r.GET("/articles/", getAllArticles)
	r.POST("/articles/", authRequired, addArticle)

	r.GET("/comment/get/:articleId", authRequired, getArticleComments)
	r.POST("/comment/like/:commentId", addLike)
	r.POST("/comment/dislike/:commentId", addDislike)
	r.POST("/comment", addComment)
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Note: do not use * in production
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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
