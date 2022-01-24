package server

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sources/data"
	"strconv"
)

type EntryJsonBody struct {
	Title string `form:"title" json:"title"`
	Text  string `form:"text" json:"text"`
}

type CommentJsonBody struct {
	EntryId uint   `form:"entryId" json:"entryId"`
	Text    string `form:"text" json:"text"`
}

func getAllArticles(c *gin.Context) {
	entries := data.GetAllArticles()
	c.JSON(http.StatusOK, entries)
}

func getFullArticleData(c *gin.Context) {
	entryId := c.Param("entryId")
	entry, comments := data.GetFullArticleData(entryId)
	c.JSON(http.StatusOK, fullEntry{Entry: entry, Comments: comments})
}

func getArticlesForPage(c *gin.Context) {
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
	entries := data.GetArticlesPerPage(page - 1)
	c.JSON(http.StatusOK, entries)
}

func addArticle(c *gin.Context) {
	username := sessions.Default(c).Get(userkey).(string)
	//user := GetUser(username)
	var requestBody EntryJsonBody

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("%s\n", err)
	}

	fmt.Println(requestBody.Text)

	data.AddArticle(username, requestBody.Title, requestBody.Text)
}

func getArticleComments(c *gin.Context) {
	//entryId, _ := strconv.ParseUint(c.Param("entryId"), 10, 32)
	entryId := c.Param("entryId")
	c.JSON(http.StatusOK, data.GetArticleComments(entryId))
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

	data.AddComment(requestBody.EntryId, username, requestBody.Text)
}

func getComments(c *gin.Context) {
	c.JSON(http.StatusOK, data.GetAllComments())
}
