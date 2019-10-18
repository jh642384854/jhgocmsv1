package advs

import (
	"github.com/gin-gonic/gin"
	"jhgocms/constant"
	"jhgocms/model"
	"jhgocms/serializer"
	"jhgocms/service/admin"
	"log"
	"net/http"
	"strconv"
)

//下面只是定义一个变量，但是还需要赋值后才能进行操作
var (
	advs *model.Advs
	advService = new(admin.AdvsService)
	msg serializer.Msg
)

//获取所有广告类别
func GetAllCategories(c *gin.Context)  {
	page, _ := strconv.Atoi(c.Query("page"))
	//limit, _ := strconv.Atoi(c.Query("limit"))
	condition := make(map[string]interface{})
	advCates,err := advService.GetAllAdvCates(condition,page,30)
	total ,_ := advService.AllRecords(model.Advcates{},condition)
	if err != nil{
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
	}else{
		msg = serializer.BuildListResponse(advCates,total)
	}
	c.JSON(http.StatusOK, msg)
}

//获取所有广告类型
func GetCatetypes(c *gin.Context)  {
	advCates,err := advService.GetAllAdvTypes()
	if err != nil{
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
	}else{
		msg = serializer.BuildListResponse(advCates,uint(len(advCates)))
	}
	c.JSON(http.StatusOK, msg)
}
/**
	获取所有的广告列表
 */
func GetAdvList(c *gin.Context)  {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	condition := make(map[string]interface{})
	advCates,err := advService.GetALLAdvs(condition,page,limit)
	total ,_ := advService.AllRecords(model.Advs{},condition)
	if err != nil{
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
	}else{
		msg = serializer.BuildListResponse(advCates,total)
	}
	c.JSON(http.StatusOK, msg)
}
/**
	创建广告类别
 */
func CreateAdvCate(c *gin.Context) {
	advCate := new(model.Advcates)
	if err := c.BindJSON(&advCate);err != nil{
		log.Fatal(err)
	}
	if advService.AddAdvCate(advCate){
		msg = serializer.BuildCreatedSuccessResponse(advCate.ID,advCate.CreatedAt.Format(constant.TimeFormat))
	}else{
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
	}
	c.JSON(http.StatusOK,msg)
}

/**
	修改广告类别
 */
func UpdateAdvCate(c *gin.Context) {

}

/**
	删除广告类别
 */
func DeleteAdvCate(c *gin.Context) {

}
/**
	获取单个广告类别信息
 */
func GetOneAdvCate(c *gin.Context)  {

}

/**
	创建广告类别
 */
func CreateAdv(c *gin.Context) {

}

/**
	修改广告类别
 */
func UpdateAdv(c *gin.Context) {

}

/**
	删除广告类别
 */
func DeleteAdv(c *gin.Context) {

}
/**
	获取单个广告类别信息
 */
func GetOneAdv(c *gin.Context)  {

}