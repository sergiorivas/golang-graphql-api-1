package resolvers

import (
	"test_articles/models"

	"github.com/jinzhu/gorm"
)

// CommentResolver info
type CommentResolver struct {
	DB  *gorm.DB
	Obj models.Comment
}

func (c CommentResolver) ID() int32 { return c.Obj.ID }

func (c CommentResolver) Content() string { return c.Obj.Content }

func ToCommentResolver(o models.Comment, DB *gorm.DB) CommentResolver {
	return CommentResolver{
		Obj: o,
		DB:  DB,
	}
}

func ToCommentResolverArray(objs []models.Comment, DB *gorm.DB) []CommentResolver {
	r := []CommentResolver{}
	for i := range objs {
		r = append(r, CommentResolver{
			Obj: objs[i],
			DB:  DB,
		})
	}
	return r
}
