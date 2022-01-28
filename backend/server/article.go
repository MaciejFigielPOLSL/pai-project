package server

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"sources/data"
	"strconv"
	"time"
)

// gin-swagger middleware
// swagger embed files

type ArticleJsonBody struct {
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

type FullEntryResponseSwagger struct {
	Article  ArticleResponse   `json:"article"`
	Comments []CommentResponse `json:"comments"`
}

type ArticleResponse struct {
	ArticleId  uint      `form:"ID" json:"ID"`
	AuthorId   string    `form:"AuthorId" json:"AuthorId"`
	AuthorName string    `form:"AuthorName" json:"AuthorName"`
	Title      string    `form:"Title" json:"Title"`
	Text       string    `form:"Text" json:"Text"`
	AddDate    time.Time `form:"AddDate" json:"AddDate"`
}

type CommentResponse struct {
	CommentId  uint   `form:"ID" json:"ID"`
	ArticleId  uint   `form:"ArticleId" json:"ArticleId"`
	AuthorName string `form:"AuthorName" json:"AuthorName"`
	Text       string `form:"Text" json:"Text"`
	Likes      int    `form:"Likes" json:"Likes"`
	Dislikes   int    `form:"Dislikes" json:"Dislikes"`
}

// @Summary Returns all articles in system
// @Description Returns all articles in system
// @Tags articles
// @Produce json
// @Success 200 {object} ArticleResponse
// @Router /articles/ [get]
func getAllArticles(c *gin.Context) {
	articles := data.GetAllArticles()
	fmt.Printf("Get all articles list %d\n", len(articles))
	c.JSON(http.StatusOK, articles)
}

// @Summary Returns article with comments
// @Description Returns full article data with array of its comments
// @Tags articles
// @Param        articleId   path      int  true  "Article ID"
// @Produce json
// @Success 200 {object} FullEntryResponseSwagger
// @Router /articles/:articleId [get]
func getFullArticleData(c *gin.Context) {
	articleId := c.Param("articleId")
	article, comments := data.GetFullArticleData(articleId)

	fmt.Printf("Get full article data %s\n", article)
	c.JSON(http.StatusOK, FullEntryResponse{Article: article, Comments: comments})
}

// @Summary Returns articles with pagination
// @Description Returns articles with pagination
// @Tags articles
// @Param        page   path      int  true  "Page to display"
// @Produce json
// @Success 200 {object} ArticleResponse
// @Router /articles/page/:page [get]
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

// @Summary Add article
// @Description Allows to add article to system
// @Tags articles
// @Param        article  body      ArticleJsonBody  true  "Article to add"
// @Produce json
// @Success 200
// @Router /articles/ [post]
func addArticle(c *gin.Context) {
	username := sessions.Default(c).Get(userkey).(string)
	//user := GetUser(username)
	var requestBody ArticleJsonBody

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("Add article %s by user %s\n", requestBody.Title, username)
	data.AddArticle(username, requestBody.Title, requestBody.Text)
}

// @Summary Returns article comments
// @Description Returns all comments for article
// @Tags articles
// @Param        articleId  path      int  true  "Article ID"
// @Produce json
// @Success 200 {object} CommentResponse
// @Router /comment/get/:articleId [get]
func getArticleComments(c *gin.Context) {
	//articleId, _ := strconv.ParseUint(c.Param("articleId"), 10, 32)
	articleId := c.Param("articleId")
	fmt.Printf("Get article '%s' comments\n", articleId)
	c.JSON(http.StatusOK, data.GetArticleComments(articleId))
}

// @Summary Add comment
// @Description Add comment for article
// @Tags articles
// @Param        comment  body      CommentJsonBody  true  "Comment"
// @Success 200
// @Router /comment [post]
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

// @Summary Get comments
// @Description Get call comments in system
// @Tags articles
// @Produce json
// @Success 200 {object} CommentResponse
// @Router /comment [get]
func getComments(c *gin.Context) {
	comments := data.GetAllComments()
	fmt.Printf("Get all comments\n", comments)
	c.JSON(http.StatusOK, comments)
}

// @Summary Add like
// @Description Add like to comment
// @Tags articles
// @Param        commentId  path      int  true  "Comment ID"
// @Router /comment/like/:commentId [get]
func addLike(c *gin.Context) {
	param := c.Param("commentId")
	commentId, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("Invalid commentId %s, cannot cast to int\n", param)
	}

	fmt.Printf("Add like to comment %s\n", commentId)
	data.AddLike(commentId)
}

// @Summary Add dislike
// @Description Add dislike to comment
// @Tags articles
// @Param        commentId  path      int  true  "Comment ID"
// @Router /comment/like/:commentId [get]
func addDislike(c *gin.Context) {
	param := c.Param("commentId")
	commentId, err := strconv.Atoi(param)
	if err != nil {
		fmt.Printf("Invalid commentId %s, cannot cast to int\n", param)
	}

	fmt.Printf("Add dislike to comment %s\n", commentId)
	data.AddDislike(commentId)
}
