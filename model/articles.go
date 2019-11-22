package model

import (
	"github.com/jinzhu/gorm"
	"jhgocms/utils"
	"time"
)

type Articles struct {
	ID              uint       `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedTime     int64      `json:"created_time"`
	UpdatedTime     int64      `json:"updated_time"` //这个不要定义为UpdatedAt名称，不然就会自动被应用为gorm的日期格式
	DeletedAt       *time.Time `sql:"index" json:"deleted_at"`
	Cid             uint       `json:"cid"`
	CatidPath       string     `gorm:"-" json:"catid_path"`
	Title           string     `json:"title"`
	Content         string     `json:"content"`
	Description     string     `json:"description"`
	Sort            uint       `json:"sort"`
	Views           uint       `json:"views"`
	Thumb           string     `json:"thumb"`
	GroupPictures   string     `json:"group_pictures"`
	Status          uint       `json:"status"`
	Seotitle        string     `json:"seotitle"`
	Seokeywords     string     `json:"seokeywords"`
	Seodescription  string     `json:"seodescription"`
	Tags            string     `json:"tags"`
	Attributes      string     `json:"attributes"`
	IsOutsideChain  bool       `json:"is_outside_chain"`
	OutsideChainUrl string     `json:"outside_chain_url"`
}

//自动更新updated_at字段
func (v *Articles) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("updated_time", utils.GetNowTimestamp())
	return nil
}

func (v *Articles) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("updated_time", utils.GetNowTimestamp())
	return nil
}

func (v *Articles) AfterFind(scope *gorm.Scope) {

}
