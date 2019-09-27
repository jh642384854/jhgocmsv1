package model

import "time"

type AdminUser struct {
	ID         uint       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Name       string     `json:"name"`
	Avatar     string     `json:"avatar"`
	Password   string     `json:"password"`
	Email      string     `json:"email"`
	Status     int        `json:"status"`
	LoginTimes int        `json:"login_times"`
	Roles      []int      `json:"roles"`
	Token      string     `gorm:"-"  json:"token"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" sql:"index"`
}

func (this AdminUser) NewAdminUsers() []AdminUser {
	return make([]AdminUser, 0)
}
