package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jhgocms/config"
	"jhgocms/middleware"
	"jhgocms/router"
	"log"
	"net/http"
)

var (
	r   *gin.Engine
	err error
)

/**
	参考实现  https://github.com/mydevc/go-gin-mvc
 */
func main() {
	r = gin.Default()
	// 2.中间件整合(暂时去掉，后面根据需要在添加)
	//regMiddleware(r)
	// 3.静态资源处理
	initStatic(r)
	// 4.路由注册
	router.SetRouter(r)
	// 5.启动服务
	if err = r.Run(fmt.Sprintf(":%s", config.SysConfig.Port)); err != nil {
		log.Fatal("服务启动失败", err.Error())
	}
}

/**
	初始化静态资源配置
 */
func initStatic(r *gin.Engine) {
	r.LoadHTMLGlob("dist/*.html") //这种是通过正则表达式匹配相应的模版文件。
	//还有一个函数LoadHTMLGlob(pattern string)，这个函数是传递具体哪些可用的模版文件，如下所示
	//r.LoadHTMLFiles("dist/index.html","dist/home.html")
	//静态资源映射配置
	r.Static("/static", "./dist/static") //这个函数其实内部也是调用的下面的StaticFS()这个函数
	//r.StaticFS("/static",http.Dir("./dist/static"))
	r.StaticFile("/favicon.ico", "./dist/favicon.ico")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": config.SysConfig.AppName,
		})
	})
}

/**
	注册中间件
 */
func regMiddleware(r *gin.Engine) {
	//1.整合Logrus日志中间件
	r.Use(middleware.LogrusLoggerToFileRotateMiddleware())
	//2.session组件
	r.Use(middleware.RedisSession())
}
