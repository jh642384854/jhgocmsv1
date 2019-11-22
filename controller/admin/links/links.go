package links

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"jhgocms/constant"
	admin2 "jhgocms/controller/admin"
	"jhgocms/model"
	"jhgocms/serializer"
	"jhgocms/service/admin"
	"net/http"
	"strconv"
)

var (
	links *model.Links
	linkService = new(admin.LinkService)
	msg serializer.Msg
)
//友情链接列表
func ListLinks(c *gin.Context)  {
	//测试
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	condition := make(map[string]interface{})
	links,err := linkService.GetALL(condition,page,limit)
	total ,_ := linkService.AllRecords(model.Links{},condition)
	if err != nil{
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
	}else{
		msg = serializer.BuildListResponse(links,total)
	}
	c.JSON(http.StatusOK, msg)
}

//创建友情链接
func CreateLinks(c *gin.Context)  {
	link := new(model.Links)
	if err := c.BindJSON(link);err != nil{
		admin2.BindJsonError(err.Error(),c)
		return
	}
	if linkService.Add(link){
		msg = serializer.BuildCreatedSuccessResponse(link.ID,link.CreatedAt.Format(constant.TimeFormat))
	}else{
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
	}
	c.JSON(http.StatusOK,msg)
}

//更新友情链接
func UpdateLinks(c *gin.Context)  {
	link := new(model.Links)
	if err := c.BindJSON(link);err != nil{
		admin2.BindJsonError(err.Error(),c)
		return
	}
	if linkService.Update(&model.Links{},link,link.ID){
		msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
	}else{
		msg = serializer.BuildSimpleFailResponse(constant.OperationFail)
	}
	c.JSON(http.StatusOK,msg)
}

//删除友情链接
func DeleteLinks(c *gin.Context)  {
	ids := make(map[string][]int)
	data,_ := c.GetRawData()
	if err := json.Unmarshal(data,&ids); err != nil{
		admin2.BindJsonError(err.Error(),c)
		return
	}
	if linkService.DeleteBatch(&model.Links{},ids["ids"]){
		msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
	}else{
		msg = serializer.BuildSimpleFailResponse(constant.OperationFail)
	}
	c.JSON(http.StatusOK,msg)
}

