package model

import (
	"time"
)

type Advs struct {
	ID uint `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Cid int `json:"cid"`
	Name string `json:"name"`
	ImgUrl string `json:"img_url"`
	ListOrder int `json:"list_order"`
	LinkUrl string `json:"link_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
}
