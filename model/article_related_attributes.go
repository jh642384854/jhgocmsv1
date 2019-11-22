package model

//文章属性

type ArticleRelatedAttributes struct {
	Status            []ArticleStatus     `json:"status"`
	Tags              []ArticleTags       `json:"tags"`
	Attributes        []ArticleAttributes `json:"attributes"`
	ArticleAttachment ArticleAttachment   `json:"attachment"`
}

type ArticleAttributes struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

//文章状态
type ArticleStatus struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ArticleTags struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ArticleAttachment struct {
	//允许上传图片文件类型
	ImgFileSuffix string `json:"img_file_suffix"`
	//允许上传的图片文件大小
	ImgFileSizeLimit int `json:"img_file_size_limit"`
	//允许上传的其他格式文件格式
	OtherFileSuffix string `json:"other_file_suffix"`
	//允许上传的其他格式文件大小
	OtherFileSizeLimit int `json:"other_file_size_limit"`
	//允许批量上传的文件大小
	BatchFileSize int `json:"batch_file_size"`
}
