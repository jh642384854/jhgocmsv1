package admin

import (
	"fmt"
	"jhgocms/model"
	"jhgocms/utils"
	"reflect"
	"strconv"
)

type ArticleService struct {
	AdminService
}

/*func (this ArticleService) GetALLArticle(condition map[string]interface{},page,limitSize int) ([]model.Articles,error){
	engine,err := xorm.NewEngine("mysql","root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	tableMapper := core.NewPrefixMapper(core.SnakeMapper{},"jh_")
	engine.SetTableMapper(tableMapper)
	var articles []model.Articles
	if err = engine.Where("id > ?",0).Find(&articles);err != nil{
		fmt.Println(err.Error())
		return nil,err
	}
	return articles,nil
}*/

func (this ArticleService) GetALLArticle(conditions map[string]model.ConditionMap,page,limitSize int) ([]model.Articles,error) {
	if limitSize == 0{
		limitSize = pageLimit
	}
	var (
		articles []model.Articles
		offset = (page-1)*limitSize
	)
	db := model.DB
	if len(conditions) > 0{
		for _, condition := range conditions {
			switch condition.Operation {
			case "=":
				v := reflect.ValueOf(condition.Value)
				kind := v.Kind()
				whereStr := ""
				switch kind {
				case reflect.String:
					whereStr = fmt.Sprintf("%s = %s",condition.Filed,condition.Value)
				case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
					whereStr = fmt.Sprintf("%s = %d",condition.Filed,condition.Value)
				}
				db = db.Where(whereStr)
			case "like":
				db = db.Where(condition.Filed + " like ?",fmt.Sprintf("%%%s%%",condition.Value))
			case "between":
				res := utils.ToSlice(condition.Value)
				begin,_ := strconv.Atoi(res[0].(string))
				end,_ := strconv.Atoi(res[1].(string))
				db = db.Where(fmt.Sprintf("%s BETWEEN %d AND %d",condition.Filed,begin,end))
			}
		}
	}
	if err := db.Limit(limitSize).Offset(offset).Order("id DESC").Find(&articles).Error; err != nil{
		return nil,err
	}
	return articles,nil
}