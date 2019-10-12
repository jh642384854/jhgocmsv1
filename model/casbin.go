package model

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"jhgocms/config"
	"strings"
)
/**
	参考实现https://github.com/it234/goapp
 */

var CasbinEnforce *casbin.Enforcer
func initCasbinEnforcer() {
	a,err := gormadapter.NewAdapterByDBUsePrefix(DB,config.SysConfig.MySql.Prefix)
	if err != nil{
		fmt.Println(a)
		fmt.Println("gormadapter.NewAdapter ",err.Error())
	}
	e,_ := casbin.NewEnforcer(config.SysConfig.Casbin.Modelfile,a)
	e.EnableLog(config.SysConfig.Casbin.Enablelog)
	/*err = e.LoadPolicy()
	if err != nil{
		fmt.Println("e.LoadPolicy ",err.Error())
	}*/
	CasbinEnforce = e
}
/**
	检查是否有权限
 */
func CheckPermission(username,path,method string) (bool, error) {
	return CasbinEnforce.Enforce(username,path,method)
}
/**
	为用户授予相应角色
 */
func AssignUserRoles(data *AdminUser)  {
	var roles []Role
	DB.Where("id in (?)",strings.Split(data.Roles,",")).Find(&roles)
	for _, role := range roles {
		CasbinEnforce.AddRoleForUser(data.Username,role.RoleName)
	}
} 
/**
	删除用户指定角色
 */
func DeleteUserRole(username string)  {
	CasbinEnforce.DeleteRolesForUser(username)
} 
/**
	为角色授予相应权限
 */
func AssignRolePermission(roleName string,data *RolePermissions)  {
	var permissions []Permissions
	DB.Select("path,method").Where("id in (?)",strings.Split(data.Permissions,",")).Find(&permissions)
	for _, permission := range permissions {
		CasbinEnforce.AddPermissionForUser(roleName,permission.Path,permission.Method)
	}
}
/**
	删除指定角色的所有权限
 */
func DeleteRolePermissions(roleName string)  {
	//删除该角色的所有权限
	CasbinEnforce.DeletePermissionsForUser(roleName)
	//删除所有管理员用户拥有角色的记录
	CasbinEnforce.DeleteRole(roleName)
}

func init() {
	//这个要优先执行一下，因为gormadapter.NewAdapterByDBUsePrefix()依赖创建的gorm.DB对象
	initMysqlEngine()
	initCasbinEnforcer()
}
