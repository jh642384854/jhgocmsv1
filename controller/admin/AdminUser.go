package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jhgocms/constant"
	"jhgocms/model"
	"jhgocms/serializer"
	"jhgocms/service/admin"
	"jhgocms/utils"
	"net/http"
	"strconv"
)
var (
	adminUser *model.AdminUser
	adminUserService *admin.AdminUserService
)
type AdminUserController struct {
	AdminBaseController
}

func Login(c *gin.Context) {
	adminuser := model.AdminUser{ID: 1, Username: "zhangsn", Avatar: "http://img.zcool.cn/community/0189495d663d4ea8012187f4e6f47c.jpg@2o.jpg", Token: "uifjdkkjkljdklfjaksfdfj",Roles:"2,7"}
	msg := serializer.BuildOneObjectResponse(adminuser)
	c.JSON(http.StatusOK, msg)
}

func UserInfo(c *gin.Context) {
	adminuser := model.AdminUser{ID: 1, Username: "zhangsn", Avatar: "http://img.zcool.cn/community/0189495d663d4ea8012187f4e6f47c.jpg@2o.jpg", Token: "uifjdkkjkljdklfjaksfdfj",Roles:"2,7"}
	msg := serializer.BuildOneObjectResponse(adminuser)
	c.JSON(http.StatusOK, msg)
}

func Logout(c *gin.Context) {
	dataMsg := serializer.CommonMsg{Status: constant.SUCCESS_TEXT, Msg: constant.LogoutSuccess}
	msg := serializer.Msg{Code: 20000, Data: dataMsg}
	c.JSON(http.StatusOK, msg)
}
/**
	获取所有后台管理用户列表
 */
func GetAllUsers(c *gin.Context)  {
	var (
		allUsers []model.AdminUser
	)
	adminUserService = new(admin.AdminUserService)
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	username := c.Query("username")
	status := c.Query("status")
	condition := make(map[string]interface{})
	if username != ""{
		userVal := make(map[string]interface{})
		userVal["like"] = username
		condition["username"] = userVal
	}
	if status != ""{
		condition["status"] = status
	}
	allUsers,err = adminUserService.GetAllUsers(condition,page,limit)
	total ,_ = adminUserService.AllRecords(model.AdminUser{},condition)
	if err != nil{
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
	}else{
		msg = serializer.BuildListResponse(allUsers,total)
	}
	c.JSON(http.StatusOK, msg)
}
/**
	创建管理员用户
 */
func AdminUserCreate(c *gin.Context)  {
	adminUser = new(model.AdminUser)
	if err = c.BindJSON(&adminUser);err != nil{
		fmt.Println(err.Error())
		msg = serializer.BuildSimpleFailResponse(constant.DataParseError)
	}else{
		adminUserService = new(admin.AdminUserService)
		adminUser.Password = utils.EncodePassword(adminUser.Password)
		msgInfo,bl := adminUserService.AddAdminUser(adminUser)
		if bl {
			msg = serializer.BuildCreatedSuccessResponse(adminUser.ID,adminUser.CreatedAt.Format(constant.TimeFormat))
		}else{
			msg = serializer.BuildSimpleFailResponse(msgInfo)
		}
	}
	c.JSON(http.StatusOK, msg)
}
/**
	更新管理员用户
 */
func AdminUserUpdate(c *gin.Context)  {
	adminUser = new(model.AdminUser)
	if err = c.BindJSON(&adminUser);err != nil{
		msg = serializer.BuildSimpleFailResponse(constant.DataParseError)
	}else{
		adminUserService = new(admin.AdminUserService)
		omitFileds := make([]string,0)
		omitFileds = append(omitFileds,"created_at","login_times")
		if adminUser.Password == ""{
			omitFileds = append(omitFileds,"password")
		}else{
			adminUser.Password = utils.EncodePassword(adminUser.Password)
		}
		msgInfo,bl := adminUserService.UpdateAdminUser(adminUser,true,omitFileds)
		if bl {
			msg = serializer.BuildSimpleSuccessResonse(msgInfo)
		}else{
			msg = serializer.BuildSimpleFailResponse(msgInfo)
		}
	}
	c.JSON(http.StatusOK, msg)
}
/**
	删除管理员用户
 */
func AdminUserDelete(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Query("id"))
	adminUserService = new(admin.AdminUserService)
	//删除管理员
	msgInfo,bl := adminUserService.DeleteAdminUser(uint(id))
	if bl {
		msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
	}else{
		msg = serializer.BuildSimpleSuccessResonse(msgInfo)
	}
	c.JSON(http.StatusOK, msg)
}


