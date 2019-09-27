package admin

import (
	"github.com/gin-gonic/gin"
	"jhgocms/config"
	"jhgocms/model"
	"jhgocms/serializer"
	"jhgocms/service/admin"
	"net/http"
	"strconv"
	"strings"
)

var (
	roleService *admin.RolesService
	role *model.Role
)

func GetAllRoles(c *gin.Context)  {
	var page int
	if c.Query("page") == ""{
		page = 0
	}else{
		page,_ = strconv.Atoi(c.Query("page"))
	}
	roleService = new(admin.RolesService)
	res,err := roleService.GetAllRoles(page)
	total,_ = roleService.AllRecords(model.Role{})
	if err != nil{
		msg = serializer.BuildSimpleFailResponse(config.DatabaseError)
	}else{
		msg = serializer.BuildListResponse(res,total)
	}
	c.JSON(http.StatusOK, msg)
}

/**
	创建角色
 */
func RolesCreate(c *gin.Context)  {
	role = new(model.Role)
	if err = c.Bind(&role);err != nil{
		msg = serializer.BuildSimpleFailResponse(config.DataParseError)
	}else{
		roleService = new(admin.RolesService)
		if roleService.Add(role){
			msg = serializer.BuildSimpleSuccessResonse(config.OperationSuccess)
		}else{
			msg = serializer.BuildSimpleFailResponse(config.DatabaseError)
		}
	}
	c.JSON(http.StatusOK,msg)
}

func RolesUpdate(c *gin.Context)  {

}

func RolesDelete(c *gin.Context)  {

}
/**
	获取某个角色的权限
 */
func GetPermissionsByRoleID(c *gin.Context)  {
	rid,_ := strconv.Atoi(c.Query("id"))
	role = new(model.Role)
	permissons := []string{}
	if err = model.DB.Where(rid).First(&role).Error; err != nil{
		msg = serializer.BuildSimpleFailResponse(config.DatabaseError)
	}else{
		rolePermission := new(model.RolePermissions)
		if !model.DB.Model(&role).Related(&rolePermission,"role_id").RecordNotFound(){
			permissons = strings.Split(rolePermission.Permissions,",")
		}
	}
	c.JSON(http.StatusOK,serializer.BuildListResponse(permissons,uint(cap(permissons))))
}