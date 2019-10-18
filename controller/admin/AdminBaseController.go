package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"jhgocms/controller"
	"jhgocms/serializer"
	"log"
	"net/http"
	"os"
)

var (
	err error
	msg serializer.Msg
	total uint
)

type AdminBaseController struct {
	controller.BaseController
}
/**
	简单的文件上传
 */
func SimpleImageUpload(c *gin.Context)  {
	file,header,err := c.Request.FormFile("file")
	if err != nil{
		c.String(http.StatusBadRequest,fmt.Sprintf("file error:%s",err.Error()))
		return
	}
	filename := header.Filename
	out,err := os.Create("upload/"+filename)
	if err != nil{
		log.Fatal(err)
	}
	defer out.Close()
	_,err = io.Copy(out,file)
	if err != nil{
		log.Fatal(err)
	}
	filepath := "http://localhost:8080/upload/"+filename
	msg := serializer.BuildUploadSuccessResponse(filepath)
	c.JSON(http.StatusOK,msg)
}
