package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"jhgocms/config"
	"jhgocms/constant"
	"jhgocms/controller"
	"jhgocms/model"
	"jhgocms/serializer"
	"jhgocms/service/admin"
	"jhgocms/utils"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
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
	测试邮件发送
 */
func TestSendEmail(c *gin.Context)  {
	email := new(model.Email)
	if err := c.BindJSON(email); err != nil{
		fmt.Println(err.Error())
		msg = serializer.BuildSimpleFailResponse(constant.DataParseError)
	}else{
		testEmailTitle := "测试邮件发送"
		testEmailBody := "测试邮件发送的内容"
		emailPassword := utils.AesDecrypt(email.SendPwd,config.SysConfig.Encodekey)// "vltgnhlbdwjsgehh"
		isSuccess,err := utils.SendEmail(email.SendUser,email.ReceiveEmail,testEmailTitle,testEmailBody,email.MailServer,emailPassword,email.MailPort)
		if isSuccess {
			msg = serializer.BuildSimpleSuccessResonse(constant.OperationSuccess)
		}else{
			msg =  serializer.BuildSimpleFailResponse(constant.OperationFail+err.Error())
		}
	}
	c.JSON(http.StatusOK,msg)
}

/**
	简单的文件上传
 */
func SimpleImageUpload(c *gin.Context)  {
	var (
		attachmentData *model.Attachment
		fileUoloadPath string
		filemd5val string
		dbfileMd5 string
		attachmentService *admin.AttachmentService
	)
	filemd5val = c.PostForm("filemd5")
	attachmentService = new(admin.AttachmentService)
	cid,_ := strconv.Atoi(c.PostForm("cid"))
	file,header,err := c.Request.FormFile("file")
	if err != nil{
		c.String(http.StatusBadRequest,fmt.Sprintf("file error:%s",err.Error()))
		return
	}
	oldFilename := header.Filename
	//获取文件的后缀信息
	fileSuffix := path.Ext(oldFilename)
	// TODO 还没有对上传文件的大小、类型进行判断处理

	newFileName := utils.CreateRandomString(10)+fileSuffix
	//如果一次传递多个文件的MD5值
	if strings.Index(filemd5val,",") > 0{
		// TODO 如果一次传递多个文件的MD5值 暂时没有好的办法处理 后面在完善
		dbfileMd5 = ""
		attachmentData = nil
	}else{
		attachmentData = attachmentService.GetFileByFileMd5(filemd5val)
		dbfileMd5 = filemd5val
	}
	if attachmentData == nil{
		//获取附近保存的文件路径
		uploadPath := utils.GetUploadPath()
		out,err := os.Create(uploadPath+newFileName)
		if err != nil{
			log.Fatal(err)
		}
		defer out.Close()
		_,err = io.Copy(out,file)
		if err != nil{
			log.Fatal(err)
		}
		//文件上传的全路径
		fileUoloadPath = fmt.Sprintf("%s%s:%s%s%s",config.SysConfig.Protocol,config.SysConfig.Domain,config.SysConfig.Port,strings.Replace(uploadPath,".","",-1),newFileName) //"http://localhost:8080/upload/"+filename
	}else{
		fileUoloadPath = attachmentData.FilePath
	}
	//将上传的文件信息记录到数据库中
	attachment := &model.Attachment{
		FileName:header.Filename,
		NewFileName:newFileName,
		FilePath:fileUoloadPath,
		FileSize:header.Size,
		FileMd5:dbfileMd5,
		Module:c.PostForm("module"),
		Cid:cid,
		CreatedAt:time.Now().Format(constant.TimeFormat),
		Suffix:fileSuffix[1:len(fileSuffix)],
	}
	adminService := new(admin.AdminService)
	adminService.Add(attachment)

	msg := serializer.BuildUploadSuccessResponse(fileUoloadPath)
	c.JSON(http.StatusOK,msg)
}
