package admin

import "jhgocms/model"

type LinkService struct {
	AdminService
}

/**
	获取所有广告列表
 */
func (this LinkService) GetALL(condition map[string]interface{},page,limitSize int) ([]model.Links,error) {
	if limitSize == 0{
		limitSize = pageLimit
	}
	var (
		links []model.Links
		offset = (page-1)*limitSize
	)
	db := model.DB
	db = this.doCondition(condition,db)
	if err := db.Limit(limitSize).Offset(offset).Order("id DESC").Find(&links).Error; err != nil{
		return nil,err
	}
	return links,nil
}