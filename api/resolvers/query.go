package resolvers

import (
	"test_articles/models"
)

type QueryResolver struct {
}

func (QueryResolver) Hello() string { return "Hello, world!1234" }

func (QueryResolver) Comment() CommentResolver {
	comment := models.Comment{
		ID:      123,
		Content: "Hll",
	}

	return ToCommentResolver(comment)
}

func (QueryResolver) Comments() []CommentResolver {

	comments := []models.Comment{}

	comment := models.Comment{
		ID:      123,
		Content: "Hll",
	}
	comments = append(comments, comment)

	comment = models.Comment{
		ID:      3456,
		Content: "Hllxx",
	}

	comments = append(comments, comment)

	return ToCommentResolverArray(comments)
}
