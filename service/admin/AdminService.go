package admin

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"jhgocms/constant"
	"jhgocms/model"
	"jhgocms/service"
	"reflect"
	"strconv"
	"strings"
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
	添加对象,这个有唯一性校验
 */
func (s AdminService) AddUnique(obj interface{},uniqueField,uniqueVal string) (errmsg string,res bool) {
	if model.DB.Where(uniqueField+" = ?",uniqueVal).First(obj).RecordNotFound(){
		if s.Add(obj){
			return constant.OperationSuccess,true
		}else{
			return constant.DatabaseError,false
		}
	}
	return constant.RecordRepeat,false
}
/**
	根据ID来删除记录
 */
func (s AdminService) DeleteById(obj interface{},id uint) bool {
	if err := model.DB.Delete(obj,id).Error; err != nil{
		return false
	}
	return true
}
/**
	带事务的方式进行删除操作
 */
func (s AdminService) DeleteByIdUseTrans(tx *gorm.DB,obj interface{},id uint) bool {
	if err := tx.Delete(obj,id).Error; err != nil{
		return false
	}else{
		return true
	}
}
/**
	根据ID来查询对象信息
 */
func (s AdminService) GetOneObj(obj interface{},id uint) interface{} {
	if model.DB.First(obj,id).RecordNotFound() {
		return nil
	}else{
		return obj
	}
}
/**
	获取该模型的所有记录总数
 */
func (s AdminService) AllRecords(obj interface{},condition map[string]interface{}) (uint, error)  {
	var total uint
	db := model.DB.Model(obj)
	db = s.doCondition(condition,db)
	if err := db.Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

//处理条件查询
func (s AdminService) doCondition(condition map[string]interface{},db *gorm.DB) *gorm.DB {
	for key, value := range condition {
		switch value.(type) {
		case string,int:
			if(value != ""){
				db = db.Where(key+" = ?",value)
			}
		case map[string]interface{}:
			for key2, val2 := range value.(map[string]interface{}) {
				fmt.Println(key+" "+key2,val2)
				db = db.Where(key+" "+key2+" ?",fmt.Sprintf("%%%s%%",val2))
			}
		}
	}
	return db
}

func doCondition2(condition map[string]interface{},db *gorm.DB)  {

}


/**
	对象列表(待完善)
	想用通用的一个方法来适配所有的模型，但是现在能力还有点欠缺，使用反射弄了半天，都没有弄好
 */
func (s AdminService) List(condition interface{},objects interface{}) interface{}{

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

/**
 *  通用更新对象。这个不考虑数据重复性的问题
 * @access public
 * @param data mixed 要更新的数据
 * @param obj mixed 要更新的数据对象
 * @param id int 要更新的数据对象的主键ID
 * @return bool 返回类型
 */
func (s AdminService) Update(obj interface{},data interface{},id uint) bool  {
	if err := model.DB.Model(obj).Where("id = ?",id).Omit("created_at").Updates(data).Error; err != nil{
		return false
	}
	return true
}

/**
 *  通用删除对象。这个是逻辑删除，并不是物理删除
 * @access public
 * @param obj mixed 要更新的数据对象
 * @param id int 要更新的数据对象的主键ID
 * @return bool 返回类型
 */
func (s AdminService) Delete(obj interface{},id uint) bool {
	if err := model.DB.Where(id).Delete(obj).Error; err != nil{
		return  false
	}
	return true
}

/**
 *  通用批量删除对象。这个是逻辑删除，并不是物理删除
 * @access public
 * @param obj mixed 要更新的数据对象
 * @param ids int 要批量删除的记录ID，多个值用","隔开
 * @return bool 返回类型
 */
func (s AdminService) DeleteBatch(obj interface{},ids []int) bool {
	if(len(ids) > 0){
		idstrs := ""
		for _, value := range ids {
			idstrs += strconv.Itoa(value) + ","
		}
		if err := model.DB.Where("id in (?)",strings.TrimRight(idstrs,",")).Delete(obj).Error; err != nil{
			return  false
		}
		return true
	}
	return true
}



/**
 * 	更新对象(待完善)
 * @access public
 * @param data mixed 要更新的数据
 * @param obj mixed 要更新的数据对象
 * @param unique bool 是否进行唯一性校验
 * @param id int 要更新的当前记录ID
 * @param uniqueFiled string 唯一性验证的字段
 * @param uniqueFiledVal string 唯一性验证的字段的字段值
 * @param omitFileds string 省略更新的字段
 * @return array 返回类型
 */
func (s AdminService) Update2(data interface{},obj interface{},unique bool,id uint,uniqueFiled string,uniqueFiledVal string,omitFileds string) (errmsg string,res bool) {
	//唯一性判断
	if unique {
		if !model.DB.Where("id != ? AND "+uniqueFiled+" = ?",id,uniqueFiledVal).First(&obj).RecordNotFound(){
			return  constant.RecordRepeat,false
		}
	}
	if err := model.DB.Model(&obj).Omit(omitFileds).Updates(data).Error;err != nil{
		return constant.DatabaseError,false
	}
	return constant.OperationSuccess,true
}
