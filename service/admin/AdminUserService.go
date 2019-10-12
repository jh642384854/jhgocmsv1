package admin

import (
	"jhgocms/constant"
	"jhgocms/model"
	"strings"
)

type AdminUserService struct {
	AdminService
}
/**
	获取所有的用户信息
 */
func (this AdminUserService) GetAllUsers(condition map[string]interface{},page,limitSize int) (users []model.AdminUser,err error) {
	var (
		adminusers []model.AdminUser
		offset = (page-1)*limitSize
	)
	db := model.DB
	db = this.doCondition(condition,db)
	if err := db.Limit(limitSize).Offset(offset).Order("id DESC").Find(&adminusers).Error; err != nil{
		return nil,err
	}
	return adminusers,nil
}
/**
	添加管理员用户信息
	@param data 是要更新的用户信息
 */
func (this AdminUserService) AddAdminUser(data *model.AdminUser) (errmsg string,res bool) {
	adminUser := new(model.AdminUser)
	if model.DB.Where("username = ?",data.Username).First(&adminUser).RecordNotFound(){
		if this.Add(data){
			model.AssignUserRoles(data)
			return constant.OperationSuccess,true
		}else{
			return constant.DatabaseError,false
		}
	}
	return  constant.RecordRepeat,false
}

/**
	更新管理员用户信息
	@param data 是要更新的用户信息
	@param unique 是否要做数据唯一性的校验
	@param omitFileds 忽略那些更新的字段
 */
func (this AdminUserService) UpdateAdminUser(data *model.AdminUser,unique bool,omitFileds []string) (errmsg string,res bool) {
	adminUser := new(model.AdminUser)
	if unique {
		if !model.DB.Where("id != ? AND username = ?",data.ID,data.Username).First(&adminUser).RecordNotFound(){
			return  constant.RecordRepeat,false
		}
	}
	if err := model.DB.Model(adminUser).Omit(strings.Join(omitFileds,",")).Updates(data).Error;err != nil{
		return constant.DatabaseError,false
	}
	model.AssignUserRoles(data)
	return constant.OperationSuccess,true
}

/**
	删除管理员
 */
func (this AdminUserService) DeleteAdminUser(id uint) (errmsg string,res bool) {
	adminUser := new(model.AdminUser)
	if model.DB.Find(adminUser,id).RecordNotFound(){
		return constant.NoRecord,false
	}else{
		if !this.DeleteById(adminUser,id){
			return constant.DatabaseError,false
		}
		//删除用户之前被设置的角色信息
		model.DeleteUserRole(adminUser.Username)
		return constant.OperationSuccess,true
	}
}