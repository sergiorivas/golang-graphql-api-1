package resolvers

import (
	"context"
	"test_articles/models"
)

func (RootResolver) Hello() string { return "Hello, world!1234" }

func (q RootResolver) FirstComment() CommentResolver {
	var c models.Comment
	q.DB.First(&c)
	return ToCommentResolver(c, q.DB)
}

func (q RootResolver) Comments() []CommentResolver {
	var comments []models.Comment
	q.DB.Find(&comments)
	return ToCommentResolverArray(comments, q.DB)
}

func (q RootResolver) Article(ctx context.Context, args struct{ ID int32 }) ArticleResolver {
	var a models.Article
	q.DB.First(&a, args.ID)
	return ToArticleResolver(a, q.DB)
}

func (q RootResolver) Articles() []ArticleResolver {
	var articles []models.Article
	q.DB.Find(&articles)
	return ToArticleResolverArray(articles, q.DB)
}
