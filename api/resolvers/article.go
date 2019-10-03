package resolvers

import (
	"test_articles/models"

	"github.com/jinzhu/gorm"
)

type ArticleResolver struct {
	DB  *gorm.DB
	Obj models.Article
}

func (r ArticleResolver) ID() int32 { return r.Obj.ID }

func (r ArticleResolver) Title() string { return r.Obj.Title }

func (r ArticleResolver) Comments() []CommentResolver {
	var comments []models.Comment

	r.DB.Model(&r.Obj).Related(&comments)
	// ANOTHER WAY 1: r.DB.Model(&r.Obj).Association("Comments").Find(&comments)
	// ANOTHER WAY 2: r.DB.Where("article_id = 1").Find(&comments)

	r.Obj.Comments = comments
	return ToCommentResolverArray(r.Obj.Comments, r.DB)
}

func ToArticleResolver(o models.Article, DB *gorm.DB) ArticleResolver {
	return ArticleResolver{
		Obj: o,
		DB:  DB,
	}
}

func ToArticleResolverArray(objs []models.Article, DB *gorm.DB) []ArticleResolver {
	r := []ArticleResolver{}
	for i := range objs {
		r = append(r, ArticleResolver{
			Obj: objs[i],
			DB:  DB,
		})
	}
	return r
}
