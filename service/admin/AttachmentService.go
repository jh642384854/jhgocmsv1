package admin

import (
	"jhgocms/model"
)

type AttachmentService struct {
	AdminService
}
/**
	根据文件的md5来获取该文件的上传记录
 */
func (this AttachmentService) GetFileByFileMd5(filemd5 string) *model.Attachment {
	data := new(model.Attachment)
	//判断多个md5值
	if model.DB.Where("file_md5 = ?",filemd5 ).First(&data).RecordNotFound(){
		return nil
	}
	return data
}