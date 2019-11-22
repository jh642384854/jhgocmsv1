package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"jhgocms/constant"
	"jhgocms/serializer"
	"jhgocms/utils"
	"net/http"
)

func BindJsonError(msg string,c *gin.Context)  {
	utils.LogrusInstance.WithFields(logrus.Fields{"error":msg}).Warn()
	msgInfo := serializer.BuildSimpleFailResponse(constant.DataParseError)
	c.JSON(http.StatusOK,msgInfo)
}