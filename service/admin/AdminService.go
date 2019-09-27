package admin

import (
	"fmt"
	"jhgocms/model"
	"jhgocms/service"
	"reflect"
)

var (
	pageLimit = 20
)

type AdminService struct {
	service.BaseService
}
/**
	添加对象
 */
func (s AdminService) Add(obj interface{}) bool {
	model.DB.Create(obj)
	return !model.DB.NewRecord(obj)
}

/**
	获取该模型的所有记录总数
 */
func (s AdminService) AllRecords(obj interface{}) (uint, error)  {
	var total uint
	if err := model.DB.Model(obj).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}
/**
	对象列表
	想用通用的一个方法来适配所有的模型，但是现在能力还有点欠缺，使用反射弄了半天，都没有弄好
 */
func (s AdminService) List(condition interface{},objects interface{}) interface{}{

	/*getValue := reflect.ValueOf(object)
	methodValue := getValue.MethodByName("NewAdminUsers")
	reflect.ValueOf(objects).Elem()
	reflect.TypeOf(objects).Name()
	methodValue.Call(nil)*/
	obj := reflect.TypeOf(objects).Name()
	var objs interface{}
	switch obj {
	case "AdminUser":
		fmt.Println("AdminUser model")
		objs = make([]model.AdminUser,0)
	}
	fmt.Println(objs)
	if err := model.DB.Limit(pageLimit).Find(&objs).Error; err != nil{
		fmt.Println(err.Error())
		return err.Error()
	}
	return objs
}