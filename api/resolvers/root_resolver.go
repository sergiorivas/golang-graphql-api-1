package resolvers

import (
	"github.com/jinzhu/gorm"
)

type RootResolver struct {
	DB *gorm.DB
}
