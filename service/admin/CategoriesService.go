package admin

import (
	"fmt"
	"strings"
	"jhgocms/constant"
	"jhgocms/model"
	"strconv"
)

type CategoriesService struct {
	AdminService
}
type PageHasNew struct {
	SubTotal uint `json:"sub_total"`
}
/**
	递归获取所有权限列表信息
 */
func (this CategoriesService) GetTreeCategories() ([]model.Categories, error) {
	Categories := make([]model.Categories, 0)
	if err := model.DB.Order("sort asc,id desc").Find(&Categories).Error; err != nil {
		return nil, err
	}
	permissonReuslt := this.recursiveCategories(Categories, 0, 0)
	return permissonReuslt, nil
}

func (this CategoriesService) recursiveCategories(Categories []model.Categories, pid, level uint) []model.Categories {
	list := []model.Categories{}
	for _, value := range Categories {
		if value.Pid == pid {
			value.Disabled = false
			value.Arrchildid = this.getArrChildid(Categories,value.ID)
			//查询是单页模型下面是否包含有新闻模型子栏目
			if value.Mid == 1{
				if strings.Index(value.Arrchildid,",") >0{
					pageHasNew := new(PageHasNew)
					model.DB.Raw("select COUNT(DISTINCT(mid)) as sub_total FROM jh_categories WHERE pid = ? or id = ?",value.ID,value.ID).Scan(&pageHasNew)
					if(pageHasNew.SubTotal <= 1){
						value.Disabled = true
					}
				}else{
					value.Disabled = true
				}
			}
			value.Level = level
			children := this.recursiveCategories(Categories, value.ID, level+1)
			if len(children)>0 {
				value.Children =children
			}
			list = append(list, value)
		}
	}
	return list
}

//获取子栏目ID列表
func (this CategoriesService) getArrChildid(Categories []model.Categories,cid uint) string {
	arrChildId := strconv.Itoa(int(cid))
	if len(Categories) > 0{
		for _, category := range Categories {
			if category.Pid>0 && category.ID != cid && category.Pid == cid{
				arrChildId += ","+this.getArrChildid(Categories,category.ID)
			}
		}
	}
	return arrChildId
}

//修正栏目的一些其他属性
func (this CategoriesService) RepairCategory()  {
	AllCategories := make([]model.Categories, 0)
	model.DB.Order("sort asc,id desc").Find(&AllCategories)
	for _, categoryData := range AllCategories {
		arrChildId := this.getArrChildid(AllCategories,categoryData.ID)
		categoryObj := new(model.Categories)
		if ! model.DB.First(categoryObj,categoryData.ID).RecordNotFound() {
			model.DB.Model(&categoryObj).Select("arrchildid").Updates(map[string]string{"arrchildid":arrChildId})
		}
	}
}

/**
	更新栏目的catid_path值
 */
func (this CategoriesService) UpdateCatCatidPath(cid uint,orginCatidPath string,isInsert bool) string {
	category := new(model.Categories)
	updateStr := ""
	if ! model.DB.First(category,cid).RecordNotFound() {
		if(orginCatidPath == "0"){
			updateStr = strconv.Itoa(int(cid))
		}else{
			if isInsert {
				updateStr = fmt.Sprintf("%s,%d",orginCatidPath,cid)
			}else{
				updateStr = orginCatidPath
			}
		}
		model.DB.Model(category).Select("catid_path").Updates(map[string]string{"catid_path":updateStr})
	}
	return updateStr
}
/**
	删除栏目
 */
func (this CategoriesService) DeleteCategory(cid uint) (bool, string) {
	article := new(model.Articles)
	//只能删除当前栏目没有数据的情况
	if model.DB.Where("cid = ?",cid).First(article).RecordNotFound(){
		if this.Delete(new(model.Categories),cid){
			return true,constant.OperationSuccess
		}else{
			return false,constant.DatabaseError
		}
	}else{
		return false,constant.CateHasSonData
	}
}