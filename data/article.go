package data

import "time"

var entriesPerPage = 10

func GetAllArticles() []Article {
	var articles []Article
	db.Order("add_date").Find(&articles)
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
		AuthorId: username,
		Title:    title,
		Text:     text,
		AddDate:  time.Now(),
	}
	db.Create(&entry)
}

func AddComment(entryId uint, authorId, text string) {
	comment := Comment{
		EntryId:  entryId,
		AuthorId: authorId,
		Text:     text,
		Likes:    0,
		Dislikes: 0,
	}
	db.Create(&comment)
}

func AddLike(commentId int32) {
	var comment Comment
	db.First(&comment, commentId)
	comment.Likes++
	db.Save(&comment)
}

func AddDislike(commentId int32) {
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
	db.Find(&comments)
	return comments
}
