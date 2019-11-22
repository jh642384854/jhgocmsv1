package model

import "time"

type Links struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Name string `json:"name"`
	Type int `json:"type"`
	ImgUrl string `json:"img_url"`
	ListOrder int `json:"list_order"`
	Status int `json:"status"`
	LinkUrl string `json:"link_url"`
	ContactUser string `json:"contact_user"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
}