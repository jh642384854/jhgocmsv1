package admin

import "jhgocms/model"

type AdminUserService struct {
	AdminService
}
/**
	获取所有的用户信息
 */
func (this AdminUserService) GetAllUsers(condition interface{},page,limitSize int) (users []model.AdminUser,err error) {
	var (
		adminusers []model.AdminUser
		offset = (page-1)*limitSize
	)
	if err := model.DB.Where(condition).Limit(limitSize).Offset(offset).Find(&adminusers).Error; err != nil{
		return nil,err
	}
	return adminusers,nil
}
