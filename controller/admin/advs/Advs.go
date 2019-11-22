package advs

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
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
	id := c.Query("id")
	if id != ""{
		if govalidator.IsNumeric(id){
			did,_ := strconv.Atoi(id)
			msgInfo,bl := advService.DeleteAdvCate(did)
			if bl {
				msg = serializer.BuildSimpleSuccessResonse(msgInfo)
			}else{
				msg = serializer.BuildSimpleFailResponse(msgInfo)
			}
		}else{
			msg = serializer.BuildSimpleFailResponse(constant.ParameterError)
		}
	}else{
		msg = serializer.BuildSimpleFailResponse(constant.MissRequiredParameter)
	}
	c.JSON(http.StatusOK,msg)
}
/**
	获取单个广告类别信息
 */
func GetOneAdvCate(c *gin.Context)  {

}

/**
	创建广告
 */
func CreateAdv(c *gin.Context) {
	adv := new(model.Advs)
	if err := c.BindJSON(adv);err != nil{
		msg = serializer.BuildSimpleFailResponse(constant.DataParseError)
	}else{
		if advService.Add(adv){
			msg = serializer.BuildCreatedSuccessResponse(uint(adv.ID),adv.CreatedAt.Format(constant.TimeFormat))
		}else{
			msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
		}
	}
	c.JSON(http.StatusOK,msg)
}

/**
	修改广告
 */
func UpdateAdv(c *gin.Context) {
	adv := new(model.Advs)
	if err := c.BindJSON(adv);err != nil{
		log.Fatal(err)
	}
	if advService.Update(&model.Advs{},adv,adv.ID){
		msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
	}else{
		msg = serializer.BuildSimpleFailResponse(constant.OperationFail)
	}
	c.JSON(http.StatusOK,msg)
}

/**
	删除广告
 */
func DeleteAdv(c *gin.Context) {
	ids := make(map[string][]int)
	data,_ := c.GetRawData()
	if err := json.Unmarshal(data,&ids); err != nil{
		log.Println(err.Error())
	}
	if advService.DeleteBatch(&model.Advs{},ids["ids"]){
		msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
	}else{
		msg = serializer.BuildSimpleFailResponse(constant.OperationFail)
	}
	c.JSON(http.StatusOK,msg)
}
/**
	获取单个广告信息
 */
func GetOneAdv(c *gin.Context)  {

}