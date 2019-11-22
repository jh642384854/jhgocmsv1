package model

import "time"

type Categories struct {
	ID          uint         `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   *time.Time   `sql:"index" json:"deleted_at"`
	Pid         uint         `json:"pid"`
	Mid         uint         `json:"mid"`
	EnName      string       `json:"en_name"`
	Name        string       `json:"name"`
	ImgUrl      string       `json:"img_url"`
	Description string       `json:"description"`
	Sort        uint         `json:"sort"`
	CatidPath   string       `json:"catid_path"`
	Arrchildid  string       `json:"arrchildid"`
	Level       uint         `gorm:"-" json:"level"`
	Disabled    bool         `gorm:"-" json:"disabled"`
	Children    []Categories `gorm:"-" json:"children"`
}
