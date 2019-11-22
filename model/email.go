package model

type Email struct {
	MailServer   string `json:"mail_server"`
	MailPort     int    `json:"mail_port"`
	SendUser     string `json:"send_user"`
	SendPwd      string `json:"send_pwd"`
	ReceiveEmail string `json:"receive_email"`
}
