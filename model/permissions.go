package model

/**
	这个模型不添加软删除的功能，所以，不需要继承BaseModel,主要就是不需要使用DeletedAt这个字段
 */
type Permissions struct {
	ID        uint          `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Method    string        `json:"method"`
	Name      string        `json:"name"`
	Path      string        `json:"path"`
	Pid       uint          `json:"pid"`
	Remark    string        `json:"remark"`
	Sort      uint          `json:"sort"`
	Level     uint          `gorm:"-" json:"level"`
	Children  []Permissions `gorm:"-" json:"children"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
}
