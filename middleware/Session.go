package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"jhgocms/config"
	"jhgocms/datasource"
	"log"
)

/**
	Session中间件
 */

func CookieSession()  {

}
/**
	使用redis连接池创建的session拦截器
 */
func RedisSession() gin.HandlerFunc {
	session_name := config.SysConfig.Redis.SessionName
	store,err := redis.NewStoreWithPool(datasource.RedisPool(),[]byte(config.SysConfig.Redis.SessionSecret))
	store.Options(sessions.Options{config.SysConfig.Redis.Path,config.SysConfig.Redis.Domain,config.SysConfig.Redis.MaxAge,config.SysConfig.Redis.Secure,config.SysConfig.Redis.HttpOnly})
	if err != nil{
		log.Println(err.Error())
	}
	return func(c *gin.Context) {
		sessions.Sessions(session_name,store)
	}
}