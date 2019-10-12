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
)

var (
	permission        *model.Permissions
	permissionService *admin.PermissionsService
	hasExits          bool
)
/**
	创建权限
 */
func PromissionCreate(c *gin.Context) {
	permission = new(model.Permissions)
	if err = c.BindJSON(&permission); err != nil {
		fmt.Println(err.Error())
	}
	permissionService = new(admin.PermissionsService)
	hasExits = permissionService.FindPermissionByNameAndPath(permission.Name, permission.Path)
	if !hasExits {
		b := permissionService.AddPermission(permission)
		if b {
			msg = serializer.BuildCreatedSuccessResponse(permission.ID, permission.CreatedAt) //.Format(config.TimeFormat)
		} else {
			msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
		}
	} else {
		msg = serializer.BuildSimpleFailResponse(constant.RecordRepeat)
	}
	c.JSON(http.StatusOK, msg)
}

/**
	更新权限
 */
func PromissionUpdate(c *gin.Context) {
	permission = new(model.Permissions)
	if err = c.BindJSON(&permission); err != nil {
		fmt.Println(err.Error())
		return
	}
	permissionService = new(admin.PermissionsService)
	res, errMsg := permissionService.UpdatePermission(permission)
	if res {
		msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
	} else {
		msg = serializer.BuildSimpleFailResponse(errMsg)
	}
	c.JSON(http.StatusOK, msg)
}

/**
	删除权限
 */
func PromissionDelete(c *gin.Context) {
	did, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		msg = serializer.BuildSimpleFailResponse(constant.ParameterError)
	}
	permissionService = new(admin.PermissionsService)
	result, err := permissionService.DeleteOnePermission(did)
	if err != nil {
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
	}
	if result {
		msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
	} else {
		msg = serializer.BuildSimpleFailResponse(constant.OperationFail)
	}
	c.JSON(http.StatusOK, msg)
}

/**
	获取权限列表
 */
func PromissionGetTreeList(c *gin.Context) {
	var (
		allPermission []model.Permissions
		allTotal      uint
		err           error
		dbop1         = true
		dbop2         = true
	)
	permissionService = new(admin.PermissionsService)
	allPermission, err = permissionService.GetTreePermission()
	if err != nil {
		dbop1 = false
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
	}
	allTotal, err = permissionService.GetTotalPermission()
	if err != nil {
		dbop2 = false
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
	}
	if dbop1 && dbop2 {
		msg = serializer.BuildListResponse(allPermission, allTotal)
	}
	c.JSON(http.StatusOK, msg)
}
