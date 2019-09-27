package admin

import "jhgocms/model"

type RolesService struct {
	AdminService
}

/**
	获取所有的用户信息
 */
func (this RolesService) GetAllRoles(page int) ([]model.Role,error) {
	var (
		rolesResult []model.Role
		offset = (page-1)*pageLimit
	)
	if err := model.DB.Limit(pageLimit).Offset(offset).Find(&rolesResult).Error; err != nil{
		return nil,err
	}
	return rolesResult,nil
}


