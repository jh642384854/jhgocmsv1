package admin

import (
	"github.com/jinzhu/gorm"
	"jhgocms/constant"
	"jhgocms/model"
	"strings"
)

type RolesService struct {
	AdminService
}

/**
	获取所有的用户信息
 */
func (this RolesService) GetAllRoles(page int) ([]model.Role,error) {
	var (
		rolesResult []model.Role
		offset = (page-1)*pageLimit
	)
	if err := model.DB.Limit(pageLimit).Offset(offset).Order("id DESC").Find(&rolesResult).Error; err != nil{
		return nil,err
	}
	return rolesResult,nil
}
/**
	更新角色信息
 */
func (this RolesService) UpdateRole(data *model.Role,unique bool,omitFileds []string) (errmsg string,res bool) {
	role := new(model.Role)
	if unique {
		if !model.DB.Where("id != ? AND role_name = ?",data.ID,data.RoleName).First(&role).RecordNotFound(){
			return  constant.RecordRepeat,false
		}
	}
	omitFiledsStr := ""
	if omitFileds != nil{
		omitFiledsStr = strings.Join(omitFileds,",")
	}
	if err := model.DB.Model(role).Omit(omitFiledsStr).Updates(data).Error;err != nil{
		return constant.DatabaseError,false
	}
	return constant.OperationSuccess,true
}

/**
	为指定角色分配相应的权限
 */
func (this RolesService) AssignRolePermission(rolePermission *model.RolePermissions) (errmsg string,res bool){
	//判断记录是否存在，不存在记录就要进行添加操作，存在记录就进行更新操作。
	rolePermissionModel := new(model.RolePermissions)
	roleModel := new(model.Role)
	if model.DB.First(roleModel,rolePermission.RoleID).RecordNotFound(){
		return constant.NoRecord,false
	}else{
		//删除当前角色之前配置的所有权限
		model.DeleteRolePermissions(roleModel.RoleName)
		if model.DB.Where("role_id = ?",rolePermission.RoleID).First(rolePermissionModel).RecordNotFound() {
			if this.Add(rolePermission){
				model.AssignRolePermission(roleModel.RoleName,rolePermission)
				return constant.OperationSuccess,true
			}else{
				return constant.DatabaseError,false
			}
		}else{
			if err := model.DB.Model(rolePermission).Select("permissions").Where("role_id = ?",rolePermission.RoleID).Updates(rolePermission).Error; err != nil{
				return constant.DatabaseError,false
			}else{
				model.AssignRolePermission(roleModel.RoleName,rolePermission)
				return constant.OperationSuccess,true
			}
		}
	}

}

/**
	删除角色相应的权限
 */
func (this RolesService) DeleteRolePermission(tx *gorm.DB,role_id uint) bool {
	if err := tx.Where("role_id = ?",role_id).Delete(model.RolePermissions{}).Error; err != nil{
		return false
	}
	return true
}
