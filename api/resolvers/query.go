package resolvers

import (
	"context"
	"errors"
	"test_articles/middleware"
	"test_articles/models"
)

func (RootResolver) Hello() string { return "Hello, world!1234" }

func (RootResolver) FirstComment(ctx context.Context) (CommentResolver, error) {
	db := middleware.GetConnection(ctx)
	var c models.Comment
	db.First(&c)
	return ToCommentResolver(c), nil
}

func (RootResolver) Comments(ctx context.Context) ([]CommentResolver, error) {
	db := middleware.GetConnection(ctx)
	var comments []models.Comment
	db.Find(&comments)
	return ToCommentResolverArray(comments), nil
}

func (RootResolver) Article(ctx context.Context, args struct{ ID int32 }) (ArticleResolver, error) {
	db := middleware.GetConnection(ctx)
	var a models.Article
	db.First(&a, args.ID)
	if a.ID == 0 {
		return ArticleResolver{}, errors.New("ID not found")
	}
	return ToArticleResolver(a), nil
}

func (RootResolver) Articles(ctx context.Context) ([]ArticleResolver, error) {
	db := middleware.GetConnection(ctx)
	var articles []models.Article
	db.Find(&articles)
	return ToArticleResolverArray(articles), nil
}
