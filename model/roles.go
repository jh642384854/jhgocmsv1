package model

type Role struct {
	BaseModel
	RoleName string `json:"role_name"`
	Sort     int    `json:"sort"`
	Remark   string `json:"remark"`
	RolePermissions RolePermissions `gorm:"-" json:"role_permissions"`
}



