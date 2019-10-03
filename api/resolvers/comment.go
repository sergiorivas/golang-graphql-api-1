package resolvers

import (
	"test_articles/models"
)

type CommentResolver struct {
	Obj models.Comment
}

func ToCommentResolver(comment models.Comment) CommentResolver {
	return CommentResolver{
		Obj: comment,
	}
}

func ToCommentResolverArray(comments []models.Comment) []CommentResolver {
	r := []CommentResolver{}
	for i := range comments {
		r = append(r, CommentResolver{
			Obj: comments[i],
		})
	}
	return r
}

func (c CommentResolver) ID() int32 { return c.Obj.ID }

func (c CommentResolver) Content() string { return c.Obj.Content }
