package admin

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"jhgocms/config"
	"jhgocms/model"
	"jhgocms/serializer"
	"jhgocms/service/admin"
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
	adminuser := model.AdminUser{ID: 1, Name: "zhangsn", Avatar: "http://img.zcool.cn/community/0189495d663d4ea8012187f4e6f47c.jpg@2o.jpg", Token: "uifjdkkjkljdklfjaksfdfj"}
	msg := serializer.BuildOneObjectResponse(adminuser)
	c.JSON(http.StatusOK, msg)
}

func UserInfo(c *gin.Context) {
	adminuser := model.AdminUser{ID: 1, Name: "zhangsn", Avatar: "http://img.zcool.cn/community/0189495d663d4ea8012187f4e6f47c.jpg@2o.jpg", Token: "uifjdkkjkljdklfjaksfdfj"}
	msg := serializer.BuildOneObjectResponse(adminuser)
	c.JSON(http.StatusOK, msg)
}

func Logout(c *gin.Context) {
	dataMsg := serializer.CommonMsg{Status: config.SUCCESS_TEXT, Msg: config.LogoutSuccess}
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
		condition["username"] = username
	}
	if status != ""{
		condition["status"] = status
	}
	allUsers,err = adminUserService.GetAllUsers(condition,page,limit)
	total ,_ = adminUserService.AllRecords(model.AdminUser{})
	if err != nil{
		msg = serializer.BuildSimpleFailResponse(config.DatabaseError)
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
		msg = serializer.BuildSimpleFailResponse(config.DataParseError)
	}else{
		fmt.Println(adminUser)
		adminUserService = new(admin.AdminUserService)
		if adminUserService.Add(adminUser) {
			msg = serializer.BuildSimpleSuccessResonse(config.OperationSuccess)
		}else{
			msg = serializer.BuildSimpleFailResponse(config.DatabaseError)
		}
	}
	c.JSON(http.StatusOK, msg)
}
/**
	更新管理员用户
 */
func AdminUserUpdate(c *gin.Context)  {

}
/**
	删除管理员用户
 */
func AdminUserDelete(c *gin.Context)  {

}
