package model

type Role struct {
	BaseModel
	/*ID              uint            `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       *time.Time      `json:"deleted_at" sql:"index"`*/
	RoleName        string          `json:"role_name"`
	Sort            int             `json:"sort"`
	Remark          string          `json:"remark"`
	RolePermissions RolePermissions `gorm:"-" json:"role_permissions"`
}
