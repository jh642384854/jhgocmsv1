package serializer

import "jhgocms/config"

// 消息总结构体
type Msg struct {
	Code int64       `json:"code"`
	Data interface{} `json:"data"`
}

// 共同的消息信息
type CommonMsg struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

// 创建对象成功返回的消息对象
type CreatedMsg struct {
	CommonMsg
	LastId      uint   `json:"last_id"`
	CreatedTime string `json:"created_time"`
}

// 返回列表数据的消息对象
type ListMsg struct {
	CommonMsg
	Total uint        `json:"total"`
	Items interface{} `json:"items"`
}

// 返回某个特定对象的消息对象
type OneObjectMsg struct {
	CommonMsg
	Results interface{} `json:"results"`
}

/**
	构建一个简单的错误消息返回
 */
func BuildSimpleFailResponse(msg string) Msg {
	return Msg{
		Code: config.SuccessCode,
		Data: CommonMsg{
			Status: config.FailText,
			Msg:    msg,
		},
	}
}

/**
	构建一个简单的成功消息返回
 */
func BuildSimpleSuccessResonse(msg string) Msg {
	return Msg{
		Code: config.SuccessCode,
		Data: CommonMsg{
			Status: config.SUCCESS_TEXT,
			Msg:    msg,
		},
	}
}

/**
	构建一个创建原始成功返回
 */
func BuildCreatedSuccessResponse(insertID uint, createTime string) Msg {
	var createMsg CreatedMsg
	createMsg.LastId = insertID
	createMsg.CreatedTime = createTime
	createMsg.CommonMsg.Status = config.SUCCESS_TEXT
	createMsg.CommonMsg.Msg = config.OperationSuccess
	return Msg{
		Code: config.SuccessCode,
		Data: createMsg,
	}
}

/**
	构建一个列表数据返回
 */
func BuildListResponse(objs interface{}, total uint) Msg {
	var listMsg ListMsg
	listMsg.CommonMsg.Status = config.SUCCESS_TEXT
	listMsg.CommonMsg.Msg = config.OperationSuccess
	listMsg.Items = objs
	listMsg.Total = total
	return Msg{
		Code: config.SuccessCode,
		Data: listMsg,
	}
}

/**
	构建某个特定对象的数据返回
 */
func BuildOneObjectResponse(objs interface{}) Msg {
	//这里也显示了如何进行多层级的struct进行赋值
	var oneObjectMsg OneObjectMsg
	oneObjectMsg.CommonMsg.Status = config.SUCCESS_TEXT
	oneObjectMsg.CommonMsg.Msg = config.OperationSuccess
	oneObjectMsg.Results = objs
	return Msg{
		Code: config.SuccessCode,
		Data: oneObjectMsg,
	}
}
