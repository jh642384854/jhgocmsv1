package middleware

import (
	"jhgocms/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/**
	日志整合中间件
 */
//这个中间件生成的日志会进行每天滚动记录
func LogrusLoggerToFileRotateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//开始时间
		startTime := time.Now()
		//处理请求
		c.Next()
		//结束时间
		endTime := time.Now()
		//获取执行时间
		latencyTime := endTime.Sub(startTime)
		//获取请求方式
		reqMethod := c.Request.Method
		//获取请求路由
		reqUri := c.Request.RequestURI
		//获取状态码
		statusCode := c.Writer.Status()
		//获取请求IP
		clientIP := c.ClientIP()
		utils.LogrusInstance.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}

//将日志记录到MongoDB
func LogrusLoggerToMongoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

//将日志记录到ES
func LogrusLoggerToESMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

//将日志记录到MQ
func LogrusLoggerToMQMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}