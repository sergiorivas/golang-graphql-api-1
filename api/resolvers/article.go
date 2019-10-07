package resolvers

import (
	"context"
	"test_articles/middleware"
	"test_articles/models"
)

type ArticleResolver struct {
	Obj models.Article
}

func (r ArticleResolver) ID() int32 { return r.Obj.ID }

func (r ArticleResolver) Title() string { return r.Obj.Title }

func (r ArticleResolver) Comments(ctx context.Context) ([]CommentResolver, error) {
	db := middleware.GetConnection(ctx)
	var comments []models.Comment

	db.Model(&r.Obj).Related(&comments)
	// ANOTHER WAY 1: db.Model(&r.Obj).Association("Comments").Find(&comments)
	// ANOTHER WAY 2: db.Where("article_id = 1").Find(&comments)

	r.Obj.Comments = comments
	return ToCommentResolverArray(r.Obj.Comments), nil
}

func ToArticleResolver(o models.Article) ArticleResolver {
	return ArticleResolver{
		Obj: o,
	}
}

func ToArticleResolverArray(objs []models.Article) []ArticleResolver {
	r := []ArticleResolver{}
	for i := range objs {
		r = append(r, ArticleResolver{
			Obj: objs[i],
		})
	}
	return r
}
