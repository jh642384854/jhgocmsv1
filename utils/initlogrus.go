package utils

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"jhgocms/config"
	"log"
	"os"
	"path"
	"time"
)

var LogrusInstance = logrus.New()

func initLogrusInstance() {

	logFilePath := config.SysConfig.Logger.Path
	logFileName := config.SysConfig.Logger.Name

	fmt.Println("logFilePath:"+logFilePath, "logFileName:"+logFileName)
	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//判断文件是否存在，不存在就进行创建
	_, err := os.Stat(fileName)
	if os.IsExist(err) == false {
		_, err := os.Create(fileName)
		if err != nil {
			log.Print(err.Error())
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	//实例化logrus
	//设置输出
	LogrusInstance.Out = src
	//设置日志级别
	LogrusInstance.SetLevel(logrus.WarnLevel)

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
	LogrusInstance.AddHook(lfHook)

	//LogrusInstance = Logrus
}

func init() {
	initLogrusInstance()
}
