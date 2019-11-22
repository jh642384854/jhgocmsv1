package article

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	admin2 "jhgocms/controller/admin"
	"jhgocms/constant"
	"jhgocms/model"
	"jhgocms/serializer"
	"jhgocms/service/admin"
	"jhgocms/utils"
	"net/http"
	"strconv"
	"strings"
)

var (
	articles       *model.Articles
	articleService = new(admin.ArticleService)
)
/**
	文章列表
 */
func List(c *gin.Context) {
	var (
		page      int    //页面数
		limit     int    //分页显示数
		cid       int    //栏目ID
		status    int    //文章状态
		title     string //模糊搜索标题
		dateRange string //创建时间范围
	)
	page, _ = strconv.Atoi(c.Query("page"))
	limit, _ = strconv.Atoi(c.Query("limit"))

	cid, _ = strconv.Atoi(c.Query("cid"))
	status, _ = strconv.Atoi(c.Query("status"))
	title = strings.Trim(c.Query("title"), "")
	dateRange = strings.Trim(c.Query("dateRange"), "")
	//封装查询条件
	condition := make(map[string]model.ConditionMap)
	if cid >0{
		condition["cid"] = model.ConditionMap{"cid","=",cid}
	}
	if status >0 {
		condition["status"] = model.ConditionMap{"status","=",status}
	}
	if title != ""{
		condition["title"] = model.ConditionMap{"title","like",title}
	}
	if dateRange != ""{
		condition["dateRange"] = model.ConditionMap{"created_time","between",strings.Split(dateRange,",")}
	}
	fmt.Println(condition)
	articles, err := articleService.GetALLArticle(condition, page, limit)
	//total, _ := articleService.AllRecords(model.Articles{}, condition)
	total := 3
	if err != nil {
		msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
	} else {
		msg = serializer.BuildListResponse(articles, uint(total))
	}
	c.JSON(http.StatusOK, msg)
}

/**
	获取单个文章信息
 */
func GetOneArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if id > 0 {
		articleObj := new(model.Articles)
		res := articleService.GetOneObj(articleObj, uint(id))
		if res == nil {
			msg = serializer.BuildSimpleFailResponse(constant.NoRecord)
		} else {
			//查询栏目的CatidPath属性
			resultData := res.(*model.Articles)
			categoriesData := new(model.Categories)
			if articleService.GetOneObj(categoriesData, resultData.Cid) != nil {
				resultData.CatidPath = categoriesData.CatidPath
			}
			msg = serializer.BuildOneObjectResponse(resultData)
		}
	} else {
		msg = serializer.BuildSimpleFailResponse(constant.MissRequiredParameter)
	}
	c.JSON(http.StatusOK, msg)
}

/**
	获取文章标签
 */
func GetArticlAttributes(c *gin.Context) {
	articleRelatedAttributes := new(model.ArticleRelatedAttributes)
	//文章状态数据
	articleStatus := []model.ArticleStatus{}
	articleStatus1 := model.ArticleStatus{1, "显示", 1}
	articleStatus2 := model.ArticleStatus{2, "不显示", 0}
	articleStatus = append(append(articleStatus, articleStatus1), articleStatus2)
	articleRelatedAttributes.Status = articleStatus

	//文章tags
	articleTags := []model.ArticleTags{}
	tag1 := model.ArticleTags{1, "tag1"}
	tag2 := model.ArticleTags{2, "tag2"}
	articleTags = append(append(articleTags, tag1), tag2)
	articleRelatedAttributes.Tags = articleTags

	//文章属性
	articleAttributes := []model.ArticleAttributes{}
	attribute1 := model.ArticleAttributes{1, "属性1"}
	attribute2 := model.ArticleAttributes{2, "属性2"}
	articleAttributes = append(append(articleAttributes, attribute1), attribute2)
	articleRelatedAttributes.Attributes = articleAttributes

	//附件属性
	articleAttachment := model.ArticleAttachment{
		ImgFileSuffix:      "jpg,jpeg,png,gif,bmp",
		ImgFileSizeLimit:   512000,
		OtherFileSuffix:    "xls,xlsx,doc,docx,ppt,pptx,zip,pdf",
		OtherFileSizeLimit: 512000,
		BatchFileSize:      3,
	}
	articleRelatedAttributes.ArticleAttachment = articleAttachment
	msg = serializer.BuildOneObjectResponse(articleRelatedAttributes)
	c.JSON(http.StatusOK, msg)
}

/**
	添加文章
 */
func AddArticle(c *gin.Context) {
	article := new(model.Articles)
	if err := c.BindJSON(article); err != nil {
		fmt.Println(err.Error())
		msg = serializer.BuildSimpleFailResponse(constant.DataParseError)
	} else {
		if articleService.Add(article) {
			msg = serializer.BuildCreatedSuccessResponse(article.ID, utils.ConverTimesstampToDatetime(article.CreatedTime))
		} else {
			msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
		}
	}
	c.JSON(http.StatusOK, msg)
}

/**
	修改文章
 */
func UpdateArticle(c *gin.Context) {
	articleData := new(model.Articles)
	if err := c.BindJSON(articleData); err != nil {
		fmt.Println(err.Error())
		msg = serializer.BuildSimpleFailResponse(constant.DataParseError)
	} else {
		if articleService.Update(new(model.Articles), articleData, articleData.ID) {
			msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
		} else {
			msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
		}
	}
	c.JSON(http.StatusOK, msg)
}

/**
	删除文章
 */
func DeleteArticle(c *gin.Context) {
	ids := make(map[string][]int)
	data,_ := c.GetRawData()
	if err := json.Unmarshal(data,&ids); err != nil{
		admin2.BindJsonError(err.Error(),c)
		return
	}
	if articleService.DeleteBatch(&model.Links{},ids["ids"]){
		msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
	}else{
		msg = serializer.BuildSimpleFailResponse(constant.OperationFail)
	}
	c.JSON(http.StatusOK,msg)
}
