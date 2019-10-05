package resolvers

import (
	"context"
	"test_articles/models"
)

type updateTitleArgs struct {
	ID    int32
	Title string
}

func (q RootResolver) UpdateTitle(ctx context.Context, args updateTitleArgs) bool {
	var article models.Article
	q.DB.Find(&article, args.ID)
	article.Title = args.Title
	q.DB.Save(&article)

	return true
}
