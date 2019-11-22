package model

import (
	"jhgocms/utils"
)

type Pages struct {
	ID        uint             `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt utils.FormatTime `json:"created_at"`
	UpdatedAt utils.FormatTime `json:"updated_at"`

	Cid         uint   `json:"cid"`
	CatName     string `gorm:"-" json:"cat_name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      uint   `json:"status"`
	Content     string `json:"content"`
}
