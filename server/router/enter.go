package router

import (
	"zga-client-manage/router/business"
	"zga-client-manage/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Autocode business.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
