package model

import "time"

type Advcates struct {
	ID          uint       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index"`
	Advs        int        `gorm:"-" json:"advs"`
	Name        string     `json:"name"`
	Type        int        `json:"type"`
	Width       int        `json:"width"`
	Height      int        `json:"height"`
	Description string     `json:"description"`
}
