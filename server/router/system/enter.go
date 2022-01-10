package system

type RouterGroup struct {
	ApiRouter
	JwtRouter
	BaseRouter
	MenuRouter
	UserRouter
	CasbinRouter
	AuthorityRouter
	OperationRecordRouter
}
