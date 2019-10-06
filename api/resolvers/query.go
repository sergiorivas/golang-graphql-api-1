package resolvers

import (
	"context"
	"errors"
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

func (q RootResolver) Article(ctx context.Context, args struct{ ID int32 }) (ArticleResolver, error) {
	var a models.Article
	q.DB.First(&a, args.ID)
	if a.ID == 0 {
		return ArticleResolver{}, errors.New("ID not found")
	}
	return ToArticleResolver(a, q.DB), nil
}

func (q RootResolver) Articles() []ArticleResolver {
	var articles []models.Article
	q.DB.Find(&articles)
	return ToArticleResolverArray(articles, q.DB)
}
