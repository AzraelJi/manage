package system

type ServiceGroup struct {
	JwtService
	MenuService
	UserService
	CasbinService
	BaseMenuService
	AuthorityService
	OperationRecordService
}
