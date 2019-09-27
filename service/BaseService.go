package service

import "github.com/jinzhu/gorm"

type BaseService struct {
	Db gorm.DB
}

