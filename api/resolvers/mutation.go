package resolvers

import (
	"context"
	"test_articles/middleware"
	"test_articles/models"
)

type updateTitleArgs struct {
	ID    int32
	Title string
}

func (RootResolver) UpdateTitle(ctx context.Context, args updateTitleArgs) (bool, error) {
	db := middleware.GetConnection(ctx)
	var article models.Article
	db.Find(&article, args.ID)
	article.Title = args.Title
	db.Save(&article)
	return true, nil
}
