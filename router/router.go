package router

import (
	"github.com/gin-gonic/gin"
	"jhgocms/controller/admin"
	"jhgocms/controller/admin/advs"
	"jhgocms/controller/admin/article"
	"jhgocms/controller/admin/links"
	"jhgocms/controller/admin/sysconfig"
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
		adminGroup.POST("upload",admin.SimpleImageUpload)

		//管理员相关操作
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
		//权限相关操作
		promissions := adminGroup.Group("promissions")
		{
			promissions.POST("create", admin.PromissionCreate)
			promissions.POST("update",admin.PromissionUpdate)
			promissions.GET("delete", admin.PromissionDelete)
			promissions.GET("treelist", admin.PromissionGetTreeList)
			promissions.GET("generateRoutes", admin.GenerateRoutesByRoles)
		}
		//角色相关操作
		role := adminGroup.Group("roles")
		{
			role.GET("list",admin.GetAllRoles)
			role.POST("create", admin.RolesCreate)
			role.POST("update",admin.RolesUpdate)
			role.GET("delete", admin.RolesDelete)
			role.GET("getpermissions", admin.GetPermissionsByRoleID)
			role.POST("savepermission",admin.SavePermission)
		}
		//广告相关操作
		advsRouter := adminGroup.Group("advs")
		{
			advsRouter.GET("categories",advs.GetAllCategories)
			advsRouter.POST("createcate",advs.CreateAdvCate)
			advsRouter.POST("updatecate",advs.UpdateAdvCate)
			advsRouter.GET("deletecate",advs.DeleteAdvCate)
			advsRouter.GET("getonecate",advs.GetOneAdvCate)

			advsRouter.GET("catetypes",advs.GetCatetypes)

			advsRouter.GET("list",advs.GetAdvList)
			advsRouter.POST("create",advs.CreateAdv)
			advsRouter.POST("update",advs.UpdateAdv)
			advsRouter.POST("delete",advs.DeleteAdv)
			advsRouter.GET("getone",advs.GetOneAdv)
		}
		//友情链接操作
		linksRouter := adminGroup.Group("links")
		{
			linksRouter.POST("create",links.CreateLinks)
			linksRouter.GET("list",links.ListLinks)
			linksRouter.POST("update",links.UpdateLinks)
			linksRouter.POST("delete",links.DeleteLinks)
		}

		//系统设置
		settings := adminGroup.Group("settings")
		{
			settings.GET("sys", sysconfig.GetConfig)
			settings.GET("customvariables", sysconfig.GetConfig)
			settings.GET("email", sysconfig.GetConfig)
			settings.POST("sys", sysconfig.SaveConfig)
			settings.POST("customvariables", sysconfig.SaveConfig)
			settings.POST("email", sysconfig.SaveConfig)
			settings.POST("testsendemail", admin.TestSendEmail)
		}

		//内容管理
		content := adminGroup.Group("articles")
		{
			content.GET("list",article.List)
			content.GET("getone/:id",article.GetOneArticle)
			content.GET("attributes",article.GetArticlAttributes)
			content.POST("create",article.AddArticle)
			content.POST("update",article.UpdateArticle)
			content.POST("delete",article.DeleteArticle)
		}

		//文章栏目管理
		categories := adminGroup.Group("categories")
		{
			categories.GET("list",article.CategoriesList)
			categories.GET("singlelist",article.PageList)
			categories.GET("delete",article.DeleteCategory)
			categories.POST("create",article.AddCategory)
			categories.POST("update",article.UpdateCategory)
		}
	}
}