package router

import (
	"github.com/gin-gonic/gin"
	"jhgocms/controller/admin"
)

const (
	adminPrefix = "admin"
	cateproxyPrefix = "category"
	shoppintPrefix = "shopping"
	v1Prefix = "v1"
	SUCCESS_TEXT = "success"
	FAIL_TEXT = "fail"
)
/**
	session配置
	crsf配置
 */
func init() {

}
func SetRouter(r *gin.Engine)  {
	var (
		adminGroup *gin.RouterGroup
	)
	adminGroup = r.Group("admin")
	{
		adminuser := adminGroup.Group("adminuser")
		{
			adminuser.POST("login", admin.Login)
			adminuser.GET("info",admin.UserInfo)
			adminuser.POST("logout", admin.Logout)
			adminuser.GET("list",admin.GetAllUsers)
			adminuser.POST("create", admin.AdminUserCreate)
			adminuser.POST("update",admin.AdminUserUpdate)
			adminuser.GET("delete", admin.AdminUserDelete)
		}
		promissions := adminGroup.Group("promissions")
		{
			promissions.POST("create", admin.PromissionCreate)
			promissions.POST("update",admin.PromissionUpdate)
			promissions.GET("delete", admin.PromissionDelete)
			promissions.GET("treelist", admin.PromissionGetTreeList)
			promissions.GET("generateRoutes", admin.GenerateRoutesByRoles)
		}
		role := adminGroup.Group("roles")
		{
			role.GET("list",admin.GetAllRoles)
			role.POST("create", admin.RolesCreate)
			role.POST("update",admin.RolesUpdate)
			role.GET("delete", admin.RolesDelete)
			role.GET("getpermissions", admin.GetPermissionsByRoleID)
			role.POST("savepermission",admin.SavePermission)
		}
	}

}