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
	ArticleId  uint   `form:"articleId" json:"articleId"`
	AuthorName string `form:"authorName" json:"authorName"`
	Text       string `form:"text" json:"text"`
}

type FullEntryResponse struct {
	Article  data.Article   `json:"article"`
	Comments []data.Comment `json:"comments"`
}

func getAllArticles(c *gin.Context) {
	articles := data.GetAllArticles()
	fmt.Printf("Get all articles list %d\n", len(articles))
	c.JSON(http.StatusOK, articles)
}

func getFullArticleData(c *gin.Context) {
	articleId := c.Param("articleId")
	article, comments := data.GetFullArticleData(articleId)

	fmt.Printf("Get full article data %s\n", article)
	c.JSON(http.StatusOK, FullEntryResponse{Article: article, Comments: comments})
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

	fmt.Printf("Get articles for page %d\n", page)
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

	fmt.Printf("Add article %s by user %s\n", requestBody.Title, username)
	data.AddArticle(username, requestBody.Title, requestBody.Text)
}

func getArticleComments(c *gin.Context) {
	//articleId, _ := strconv.ParseUint(c.Param("articleId"), 10, 32)
	articleId := c.Param("articleId")
	fmt.Printf("Get article '%s' comments\n", articleId)
	c.JSON(http.StatusOK, data.GetArticleComments(articleId))
}

func addComment(c *gin.Context) {
	//username := sessions.Default(c).Get(userkey).(string)
	//user := GetUser(username)
	var requestBody CommentJsonBody

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("%d: %s, %s\n", requestBody.ArticleId, requestBody.AuthorName, requestBody.Text)

	data.AddComment(requestBody.ArticleId, requestBody.AuthorName, requestBody.Text)
}

func getComments(c *gin.Context) {
	comments := data.GetAllComments()
	fmt.Printf("Get all comments\n", comments)
	c.JSON(http.StatusOK, comments)
}

func addLike(c *gin.Context) {
	param := c.Param("commentId")
	commentId, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("Invalid commentId %s, cannot cast to int\n", param)
	}

	fmt.Printf("Add like to comment %s\n", commentId)
	data.AddLike(commentId)
}

func addDislike(c *gin.Context) {
	param := c.Param("commentId")
	commentId, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("Invalid commentId %s, cannot cast to int\n", param)
	}

	fmt.Printf("Add dislike to comment %s\n", commentId)
	data.AddDislike(commentId)
}
