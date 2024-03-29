package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"jhgocms/constant"
	"time"
)

type ConditionMap struct {
	Filed string
	Operation string
	Value interface{}
}
/**
	参考 https://www.cnblogs.com/mrylong/p/11326792.html
 */
type BaseModel struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (m BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_at", NowTime())
	scope.SetColumn("updated_at", NowTime())
	return nil
}

/*func (m BaseModel) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("updated_at", NowTime())
	return nil
}*/

func NowTime() string {
	return time.Now().Format(constant.TimeFormat)
}

func AddWhere(field,operation string,value interface{}) *gorm.DB {
	db := DB.Where(fmt.Sprintf("%s %s %s",field,operation,value))
	return db
}