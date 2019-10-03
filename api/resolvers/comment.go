package resolvers

import (
	"test_articles/models"
)

type CommentResolver struct {
	Obj models.Comment
}

func (c CommentResolver) ID() int32 { return c.Obj.ID }

func (c CommentResolver) Content() string { return c.Obj.Content }

func ToCommentResolver(o models.Comment) CommentResolver {
	return CommentResolver{
		Obj: o,
	}
}

func ToCommentResolverArray(objs []models.Comment) []CommentResolver {
	r := []CommentResolver{}
	for i := range objs {
		r = append(r, CommentResolver{
			Obj: objs[i],
		})
	}
	return r
}
