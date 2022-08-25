package WebsiteStorage

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
}

func NewSQL(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
