package config

import (
	"fmt"
	"github.com/spf13/viper"
)

//服务端配置
type AppConfig struct {
	AppName string `json:"app_name"`
	Port    string `json:"port"`
	Domain  string `json:"domain"`
	Logger
	MySql
	Redis
	Mq
	Mongodb
}

//日志配置信息
type Logger struct {
	Path string `json:"path"` //这个地方命名为log_path就不行，不确定是为什么？还需要排查
	Name string `json:"name"`
}

//数据库配置信息
type MySql struct {
	Connstr      string `json:"connstr"`
	Drive        string `json:"drive"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Database     string `json:"database"`
	Params       string `json:"params"`
	ShowSql      bool   `json:"show_sql"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxLifetime  int    `json:"max_lifetime"`
	MaxIdleConns int    `json:"max_idle_conns"`
	Debug        bool   `json:"debug"`
	Prefix       string `json:"prefix"`
	Suffix       string `json:"suffix"`
	Charset      string `json:"charset"`
	Timeout      int    `json:"timeout"`
}

//Redis配置信息
type Redis struct {
	NetWork       string `json:"net_work"`
	Host          string `json:"host"`
	Port          string `json:"port"`
	Database      int    `json:"database"`
	Pconnect      bool   `json:"pconnect"`
	Password      string `json:"password"`
	Prefix        string `json:"prefix"`
	Timeout       int    `json:"timeout"`
	MaxIdle       int    `json:"max_idle"`
	MaxActive     int    `json:"max_active"`
	Path          string `json:"path"`
	Domain        string `json:"domain"`
	MaxAge        int    `json:"max_age"`
	Secure        bool   `json:"secure"`
	HttpOnly      bool   `json:"http_only"`
	SessionSecret string `json:"session_secret"`
	SessionName   string `json:"session_name"`
}

//MQ的相关配置
type Mq struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Vh       string `json:"vh"`
	Topic    string `json:"topic"`
	Timeout  int    `json:"timeout"`
}

//Mongodb的相关配置
type Mongodb struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

var SysConfig AppConfig
func initConfig()  {
	viper.SetConfigName("config") //这里不需要写文件的后缀名
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil{
		panic(fmt.Errorf("Fatal error config file:%s\n",err))
	}
	//下面调用viper.Unmarshal()方法，将读取的配置文件内容映射到定义的配置结构体上面
	if err = viper.Unmarshal(&SysConfig); err != nil{
		panic(fmt.Errorf("Fatal Unmarshal config:%s\n",err))
	}
}

func init() {
	initConfig()
}