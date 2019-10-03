package resolvers

import (
	"test_articles/models"

	"github.com/jinzhu/gorm"
)

type QueryResolver struct {
	DB *gorm.DB
}

func (QueryResolver) Hello() string { return "Hello, world!1234" }

func (q QueryResolver) Comment() CommentResolver {
	var c models.Comment
	q.DB.First(&c)
	return ToCommentResolver(c, q.DB)
}

func (q QueryResolver) Comments() []CommentResolver {

	var comments []models.Comment

	q.DB.Find(&comments)

	return ToCommentResolverArray(comments, q.DB)
}

func (q QueryResolver) Article() ArticleResolver {
	var a models.Article
	q.DB.First(&a)
	return ToArticleResolver(a, q.DB)
}
