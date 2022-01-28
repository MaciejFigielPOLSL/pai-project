package data

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	gorm.Model
	ArticleId  uint
	AuthorId   string
	AuthorName string
	Title      string
	Text       string
	AddDate    time.Time
}

type Comment struct {
	gorm.Model
	ArticleId  uint
	AuthorName string
	Text       string
	Likes      int
	Dislikes   int
}

var entriesPerPage = 10

func GetAllArticles() []Article {
	var articles []Article
	db.Order("add_date desc").Find(&articles)
	return articles
}

func GetArticlesPerPage(page int) []Article {
	var entries []Article
	result := db.Order("add_date").Offset(page * entriesPerPage).Limit(entriesPerPage).Find(&entries)

	if result.Error != nil {

	}
	return entries
}

func AddArticle(username string, title, text string) {
	entry := Article{
		AuthorId:   username,
		AuthorName: username,
		Title:      title,
		Text:       text,
		AddDate:    time.Now(),
	}
	db.Create(&entry)
}

func AddComment(entryId uint, authorName, text string) {
	comment := Comment{
		ArticleId:  entryId,
		AuthorName: authorName,
		Text:       text,
		Likes:      0,
		Dislikes:   0,
	}
	db.Create(&comment)
}

func AddLike(commentId int) {
	var comment Comment
	db.First(&comment, commentId)
	comment.Likes++
	db.Save(&comment)
}

func AddDislike(commentId int) {
	var comment Comment
	db.First(&comment, commentId)
	comment.Dislikes++
	db.Save(&comment)
}

func GetFullArticleData(id string) (Article, []Comment) {
	return GetArticle(id), GetArticleComments(id)
}

func GetArticle(id string) Article {
	var entry Article
	db.First(&entry, "id = ?", id)
	return entry
}

func GetArticleComments(entryId string) []Comment {
	var comments []Comment
	db.Find(&comments, "entry_id = ?", entryId)
	return comments
}

func GetArticlesForAuthor(authorId string) []Article {
	var entries []Article
	db.Where("author_id LIKE ?", authorId).Find(&entries)
	return entries
}

func GetAllComments() []Comment {
	var comments []Comment
	db.Order("created_at desc").Find(&comments)
	return comments
}

func DeleteComments() {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Comment{})
}

func DeleteArticles() {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Article{})
}
