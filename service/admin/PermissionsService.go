package admin

import (
	"jhgocms/config"
	"jhgocms/model"
)

type PermissionsService struct {
	AdminService
}

// @title 添加权限
// @version 1.0
// @description 添加成功返回true，否则返回false
// @Produce  json
// @Success bool
func (this PermissionsService) AddPermission(permission *model.Permissions) bool {
	model.DB.Create(permission)
	return !model.DB.NewRecord(permission)
}

func (this PermissionsService) UpdatePermission(permission *model.Permissions) (res bool, errMsg string) {
	hasPermission := model.DB.Where("name = ? AND path = ? AND id != ?", permission.Name, permission.Path, permission.ID).First(&model.Permissions{}).RecordNotFound()
	if !hasPermission {
		return false, config.RecordRepeat
	}
	if err := model.DB.Model(model.Permissions{}).Omit("id","created_time").Where(permission.ID).Updates(permission).Error; err != nil {
		return false, config.DatabaseError
	}
	return true, ""
}

/**
	根据权限名称和路由地址来查询是否存在该权限
 */
func (this PermissionsService) FindPermissionByNameAndPath(name, path string) bool {
	permissions := new(model.Permissions)
	model.DB.Where("name = ? AND path = ?", name, path).First(&permissions)
	if permissions.ID > 0 {
		return true
	}
	return false
}

/**
	获取权限总数
 */
func (this PermissionsService) GetTotalPermission() (uint, error) {
	var total uint
	if err := model.DB.Model(model.Permissions{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

/**
	递归获取所有权限列表信息
 */
func (this PermissionsService) GetTreePermission() ([]model.Permissions, error) {
	permissions := make([]model.Permissions, 0)
	if err := model.DB.Find(&permissions).Error; err != nil {
		return nil, err
	}
	permissonReuslt := this.recursivePermission(permissions, 0, 0)
	return permissonReuslt, nil
}

func (this PermissionsService) recursivePermission(permissions []model.Permissions, pid, level uint) []model.Permissions {
	list := []model.Permissions{}
	for _, value := range permissions {
		if value.Pid == pid {
			value.Level = level
			value.Children = this.recursivePermission(permissions, value.ID, level+1)
			list = append(list, value)
		}
	}
	return list
}

/**
	删除某条权限记录
 */
func (this PermissionsService) DeleteOnePermission(id int) (res bool, err error) {
	if err := model.DB.Delete(model.Permissions{}, "id = ?", id).Error; err != nil {
		return false, err
	}
	return true, nil
}