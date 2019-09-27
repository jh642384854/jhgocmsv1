package middleware

import (
	"jhgocms/config"
	"log"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

/**
	日志整合中间件
 */

//这个中间件生成的日志会进行每天滚动记录
func LogrusLoggerToFileRotateMiddleware() gin.HandlerFunc {

	logFilePath := config.SysConfig.Logger.Path
	logFileName := config.SysConfig.Logger.Name

	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//判断文件是否存在，不存在就进行创建
	_,err := os.Stat(fileName)
	if os.IsExist(err) == false {
		_,err :=os.Create(fileName)
		if err != nil{
			log.Print(err.Error())
			return nil
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	//实例化logrus
	logrusInstance := logrus.New()
	//设置输出
	logrusInstance.Out = src
	//设置日志级别
	logrusInstance.SetLevel(logrus.DebugLevel)

	//设置rotatelogs
	logWriter, err := rotatelogs.New(
		//分割后的文件名称
		fileName+".%Y%m%d.log",
		//生成软链接，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		//设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),
		//设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writerMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//新增Hook
	logrusInstance.AddHook(lfHook)

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

		logrusInstance.WithFields(logrus.Fields{
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