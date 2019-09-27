package model

type RolePermissions struct {
	ID int `json:"id"`
	RoleID int `json:"role_id"`
	Permissions string `json:"permissions"`
}