package model

type Attachment struct {
	ID          int `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	FileName    string `json:"file_name"`
	NewFileName string `json:"new_file_name"`
	FileMd5     string `gorm:"file_md5" json:"file_md5"`
	FileSize    int64  `json:"file_size"`
	Suffix      string `json:"suffix"`
	FilePath    string `json:"file_path"`
	Cid         int    `json:"cid"`
	Module      string `json:"module"`
	CreatedAt   string `json:"created_at"`
}
