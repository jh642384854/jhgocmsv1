package admin

import (
	"jhgocms/controller"
	"jhgocms/serializer"
)

var (
	err error
	msg serializer.Msg
	total uint
)

type AdminBaseController struct {
	controller.BaseController
}

func (this *AdminBaseController) Test1()  {

}
