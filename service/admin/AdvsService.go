package admin

import "jhgocms/model"

type AdvsService struct {
	AdminService
}
/**
	获取所有的广告类别列表
 */
func (this AdvsService) GetAllAdvCates(condition map[string]interface{},page,limitSize int) (advcates []model.Advcates,err error) {
	var (
		advCates []model.Advcates
		offset = (page-1)*limitSize
	)
	db := model.DB
	db = this.doCondition(condition,db)
	if err := db.Limit(limitSize).Offset(offset).Order("id DESC").Find(&advCates).Error; err != nil{
		return nil,err
	}
	return advCates,nil
} 
/**
	获取所有的广告类型
 */
func (this AdvsService) GetAllAdvTypes() ([]model.Advtypes,error) {
	var (
		advtypes []model.Advtypes
	)
	if err := model.DB.Find(&advtypes).Error;err != nil{
		return nil,err
	}
	return advtypes,nil
} 
/**
	获取所有广告列表
 */
func (this AdvsService) GetALLAdvs(condition map[string]interface{},page,limitSize int) ([]model.Advs,error) {
	var (
		advs []model.Advs
		offset = (page-1)*limitSize
	)
	db := model.DB
	db = this.doCondition(condition,db)
	if err := db.Limit(limitSize).Offset(offset).Order("id DESC").Find(&advs).Error; err != nil{
		return nil,err
	}
	return advs,nil
}

/**
	添加广告类别
 */
func (this AdvsService) AddAdvCate(advcates *model.Advcates) bool {
	return this.Add(advcates)
}