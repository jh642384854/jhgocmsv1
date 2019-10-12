package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AdminUser struct {
	ID         uint       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Username   string     `json:"username"`
	Avatar     string     `json:"avatar"`
	Password   string     `json:"password"`
	Email      string     `json:"email"`
	Status     int        `json:"status"`
	LoginTimes int        `json:"login_times"`
	Roles      string     `json:"roles"`
	Token      string     `gorm:"-"  json:"token"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at" sql:"index"`
}

func (this *AdminUser) NewAdminUsers() []AdminUser {
	return make([]AdminUser, 0)
}

func (this *AdminUser) AfterFind(scope *gorm.Scope)  {
	//fmt.Println("AdminUser AfterFind")
}