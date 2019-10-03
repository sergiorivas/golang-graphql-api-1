package resolvers

import (
	"test_articles/models"
)

type ArticleResolver struct {
	Obj models.Article
}

func (r ArticleResolver) ID() int32 { return r.Obj.ID }

func (r ArticleResolver) Title() string { return r.Obj.Title }

func (r ArticleResolver) Comment() *CommentResolver {
	return &CommentResolver{
		Obj: r.Obj.Comment,
	}
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
