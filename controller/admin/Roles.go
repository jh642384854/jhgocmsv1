package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jhgocms/constant"
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
/**
	获取所有的角色列表
 */
func GetAllRoles(c *gin.Context)  {
	var page int
	if c.Query("page") == ""{
		page = 0
	}else{
		page,_ = strconv.Atoi(c.Query("page"))
	}
	roleService = new(admin.RolesService)
	res,err := roleService.GetAllRoles(page)
	total,_ = roleService.AllRecords(model.Role{},nil)
	if err != nil{
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
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
		msg = serializer.BuildSimpleFailResponse(constant.DataParseError)
	}else{
		roleService = new(admin.RolesService)
		msgInfo,bl := roleService.AddUnique(role,"role_name",role.RoleName)
		if bl{
			msg = serializer.BuildCreatedSuccessResponse(role.ID,role.CreatedAt)
		}else{
			msg = serializer.BuildSimpleFailResponse(msgInfo)
		}
	}
	c.JSON(http.StatusOK,msg)
}
/**
	更新角色
 */
func RolesUpdate(c *gin.Context)  {
	role = new(model.Role)
	if err = c.Bind(&role);err != nil{
		msg = serializer.BuildSimpleFailResponse(constant.DataParseError)
	}else{
		roleService = new(admin.RolesService)
		msgInfo,bl := roleService.UpdateRole(role,true,nil)
		if bl{
			msg = serializer.BuildSimpleSuccessResonse(msgInfo)
		}else{
			msg = serializer.BuildSimpleFailResponse(msgInfo)
		}
	}
	c.JSON(http.StatusOK,msg)
}
/**
	删除角色
 */
func RolesDelete(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Query("id"))
	roleService = new(admin.RolesService)
	roleObj := new(model.Role)
	roleService.GetOneObj(roleObj,uint(id))
	//开启事务
	tx := model.DB.Begin()
	if roleService.DeleteByIdUseTrans(tx,new(model.Role),uint(id)) {
		//删除给该角色赋予的权限信息
		if !roleService.DeleteRolePermission(tx,roleObj.ID){
			msg = serializer.BuildSimpleFailResponse(constant.OperationFail)
			tx.Rollback()
		}
		model.DeleteRolePermissions(roleObj.RoleName)
		msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
	}else{
		tx.Rollback()
		msg = serializer.BuildSimpleFailResponse(constant.OperationFail)
	}
	//提交事务
	tx.Commit()
	c.JSON(http.StatusOK, msg)
}
/**
	获取某个角色的权限
 */
func GetPermissionsByRoleID(c *gin.Context)  {
	rid,_ := strconv.Atoi(c.Query("id"))
	role = new(model.Role)
	permissons := []string{}
	if err = model.DB.Where(rid).First(&role).Error; err != nil{
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
	}else{
		rolePermission := new(model.RolePermissions)
		if !model.DB.Model(&role).Related(&rolePermission,"role_id").RecordNotFound(){
			permissons = strings.Split(rolePermission.Permissions,",")
		}
	}
	c.JSON(http.StatusOK,serializer.BuildListResponse(permissons,uint(cap(permissons))))
}

/**
	为某个角色设置权限
 */
func SavePermission(c *gin.Context)  {
	rolePermission := new(model.RolePermissions)
	if err := c.BindJSON(rolePermission); err != nil{
		fmt.Println(err.Error())
		msg = serializer.BuildSimpleFailResponse(constant.DataParseError)
	}
	roleService = new(admin.RolesService)
	msgInfo,bl := roleService.AssignRolePermission(rolePermission)
	if bl{
		msg = serializer.BuildSimpleSuccessResonse(msgInfo)
	}else{
		msg = serializer.BuildSimpleFailResponse(msgInfo)
	}
	c.JSON(http.StatusOK,msg)
}