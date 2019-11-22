package admin

import (
	"jhgocms/model"
)

type SysConfigService struct {
	AdminService
}

//获取某个系统配置
func (this SysConfigService) GetConfig(key string) *model.SysConfig{
	sysconfig := new(model.SysConfig)
	if model.DB.Where("config_key = ?",key).First(&sysconfig).RecordNotFound(){
		return nil
	}else{
		return sysconfig
	}
}

//插入或更新某个类别的配置
func (this SysConfigService) SaveOrInsertCofnig(config *model.SysConfig) bool {
	sysconfig := new(model.SysConfig)
	if model.DB.Where("config_key = ?",config.ConfigKey).First(&sysconfig).RecordNotFound(){
		return this.Add(config)
	}else{
		if err := model.DB.Model(sysconfig).Omit("config_key").Updates(config).Error;err != nil{
			return false
		}else{
			return true
		}
	}

}
