package repo

import "github.com/jinzhu/gorm"

type Repo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{db}
}
