package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jhgocms/constant"
	"jhgocms/model"
	"jhgocms/serializer"
	"jhgocms/service/admin"
	"net/http"
	"strconv"
	"strings"
)

var (
	categories *model.Categories
	categoriesService = new(admin.CategoriesService)
	msg serializer.Msg
)
/**
	获取所有栏目列表
 */
func CategoriesList(c *gin.Context)  {
	var (
		allCategories  []model.Categories
	)
	allCategories,err := categoriesService.GetTreeCategories()
	if err != nil{
		fmt.Println(err.Error())
	}
	msg = serializer.BuildListResponse(allCategories, uint(len(allCategories)))

	c.JSON(http.StatusOK,msg)
}

/**
	单页栏目列表
 */
func PageList(c *gin.Context)  {
	
}

/**
	添加栏目
 */
func AddCategory(c *gin.Context)  {
	categories := new(model.Categories)
	if err := c.BindJSON(categories); err != nil{
		fmt.Println(err.Error())
		msg = serializer.BuildSimpleFailResponse(constant.DataParseError)
	}else{
		if categoriesService.Add(categories){
			resultCatidPath := categoriesService.UpdateCatCatidPath(categories.ID,categories.CatidPath,true)
			categoriesService.RepairCategory()
			msg = serializer.BuildCreatedSuccessMoreInfoResponse(categories.ID,categories.CreatedAt.Format(constant.TimeFormat),resultCatidPath)
		}else{
			msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
		}
	}
	c.JSON(http.StatusOK,msg)
}

/**
	修改栏目
 */
func UpdateCategory(c *gin.Context)  {
	categoryData := new(model.Categories)
	if err := c.BindJSON(categoryData); err != nil{
		fmt.Println(err.Error())
		msg = serializer.BuildSimpleFailResponse(constant.DataParseError)
	}else{
		//将结构体转换为map数据结构。因为对通过结构体变量更新字段值, gorm库会忽略零值字段。就是字段值等于0, nil, "", false这些值会被忽略掉，不会更新。如果想更新零值，可以使用map类型替代结构体。
		newCategoryData := make(map[string]interface{})
		newCategoryData["name"] = categoryData.Name
		newCategoryData["en_name"] = categoryData.EnName
		newCategoryData["mid"] = categoryData.Mid
		newCategoryData["pid"] = categoryData.Pid
		newCategoryData["description"] = categoryData.Description
		newCategoryData["sort"] = categoryData.Sort
		newCategoryData["img_url"] = categoryData.ImgUrl
		if categoriesService.Update(new(model.Categories),newCategoryData,categoryData.ID){
			categoriesService.UpdateCatCatidPath(categoryData.ID,categoryData.CatidPath,false)
			categoriesService.RepairCategory()
			msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
		}else{
			msg = serializer.BuildSimpleFailResponse(constant.DatabaseError)
		}
	}
	c.JSON(http.StatusOK,msg)
}

/**
	删除栏目
 */
func DeleteCategory(c *gin.Context)  {
	id := strings.TrimSpace(c.Query("id"))
	//校验参数的合法性
	if id == "" {
		msg = serializer.BuildSimpleFailResponse(constant.MissRequiredParameter)
	}else{
		id,err := strconv.Atoi(id)
		if err != nil{
			msg = serializer.BuildSimpleFailResponse(constant.ParameterError)
		}else{
			//查询文章表中是否还有关联的文章信息
			bl,info := categoriesService.DeleteCategory(uint(id))
			if bl{
				msg = serializer.BuildSimpleSuccessResonse(info)
			}else{
				msg = serializer.BuildSimpleFailResponse(info)
			}
		}
	}
	c.JSON(http.StatusOK,msg)
}