package config

import (
	"fmt"
	"github.com/spf13/viper"
)

/**
	注意：
	定义字段的时候，特别是为字段添加json的tag不能出现下划线的形式，在对应的config.json文件中定义的key也不能出现下划线的方式，比如不能写成max_open_conns，而要写成maxopenconns
 */
//服务端配置
type AppConfig struct {
	AppName string `json:"appname"`
	Port    string `json:"port"`
	Domain  string `json:"domain"`
	Logger
	MySql
	Redis
	Mq
	Mongodb
	Casbin
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
	ShowSql      bool   `json:"showsql"`
	MaxOpenConns int    `json:"maxopenconns"`
	MaxLifetime  int    `json:"maxlifetime"`
	MaxIdleConns int    `json:"maxidleconns"`
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
	MaxIdle       int    `json:"maxidle"`
	MaxActive     int    `json:"maxactive"`
	Path          string `json:"path"`
	Domain        string `json:"domain"`
	MaxAge        int    `json:"max_age"`
	Secure        bool   `json:"secure"`
	HttpOnly      bool   `json:"httponly"`
	SessionSecret string `json:"sessionsecret"`
	SessionName   string `json:"sessionname"`
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

type Casbin struct {
	Enablelog bool `json:"enablelog"`
	Modelfile string `json:"modelfile"`
}

var SysConfig AppConfig

func initConfig() {
	viper.SetConfigName("config") //这里不需要写文件的后缀名
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file:%s\n", err))
	}
	//下面调用viper.Unmarshal()方法，将读取的配置文件内容映射到定义的配置结构体上面
	if err = viper.Unmarshal(&SysConfig); err != nil {
		panic(fmt.Errorf("Fatal Unmarshal config:%s\n", err))
	}
}

func init() {
	initConfig()
}
