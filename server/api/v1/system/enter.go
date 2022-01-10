package system

import "zga-client-manage/service"

type ApiGroup struct {
	JwtApi
	BaseApi
	CasbinApi
	AuthorityApi
	AuthorityMenuApi
	OperationRecordApi
}

var (
	jwtService             = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService            = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService            = service.ServiceGroupApp.SystemServiceGroup.UserService
	casbinService          = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	baseMenuService        = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	authorityService       = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
)
