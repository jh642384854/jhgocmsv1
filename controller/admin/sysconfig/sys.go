package sysconfig

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jhgocms/constant"
	"jhgocms/model"
	"jhgocms/serializer"
	"jhgocms/service/admin"
	"net/http"
	"strings"
)
var (
	sysconfig *model.SysConfig
	sysConfigService = new(admin.SysConfigService)
	msg serializer.Msg
)
//获取配置
func GetConfig(c *gin.Context)  {
	urls := strings.Split(c.Request.URL.String(),"/")
	configKey := urls[len(urls)-1]
	sysconfig = sysConfigService.GetConfig(configKey)
	if sysconfig != nil{
		msg = serializer.BuildOneObjectResponse(sysconfig.ConfigValue)
	}else{
		msg = serializer.BuildOneObjectResponse("")
	}
	c.JSON(http.StatusOK, msg)
}
//保存配置
func SaveConfig(c *gin.Context)  {
	urls := strings.Split(c.Request.URL.String(),"/")
	configKey := urls[len(urls)-1]
	sysConfig := new(model.SysConfig)
	if err := c.BindJSON(&sysConfig);err != nil{
		fmt.Println(err.Error())
	}
	sysConfig.ConfigKey = configKey
	isSuccess := sysConfigService.SaveOrInsertCofnig(sysConfig)
	if isSuccess {
		msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
	}else{
		msg = serializer.BuildSimpleFailResponse(constant.OperationFail)
	}
	c.JSON(http.StatusOK, msg)
}