package models

type Article struct {
	ID       int32
	Title    string
	Comments []Comment `gorm:"foreignkey:ArticleId"`
}
