package helper

import (
	"gorm.io/gorm"
)

func Paginate(limit, page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func GetPage(page int) int {
	if page <= 0 {
		page = 1
	}
	return page
}

func GetLimit(limit int) int {
	if limit <= 0 {
		limit = 10
	}
	return limit
}
