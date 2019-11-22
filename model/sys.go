package model

type SysConfig struct {
	ID    uint   `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	ConfigKey   string `json:"config_key"`
	ConfigValue string `json:"config_value"`
}

/*func (this SysConfig) TableName() string {
	return "sys_config"
}*/

