package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"zga-client-manage/global"
	"zga-client-manage/middleware"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域，如需跨域可以打开下面的注释
	Router.Use(middleware.Cors())        // 直接放行全部跨域请求
	Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 获取路由组实例
	//systemRouter := router.RouterGroupApp.System
	//{
	//	systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	//}
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	//PrivateGroup := Router.Group("")
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//{
	//	systemRouter.InitApiRouter(PrivateGroup)                 // 注册功能api路由
	//	systemRouter.InitJwtRouter(PrivateGroup)                 // jwt相关路由
	//	systemRouter.InitUserRouter(PrivateGroup)                // 注册用户路由
	//	systemRouter.InitMenuRouter(PrivateGroup)                // 注册menu路由
	//	systemRouter.InitCasbinRouter(PrivateGroup)              // 权限相关路由
	//	systemRouter.InitAuthorityRouter(PrivateGroup)           // 注册角色路由
	//	systemRouter.InitSysDictionaryRouter(PrivateGroup)       // 字典管理
	//	systemRouter.InitSysOperationRecordRouter(PrivateGroup)  // 操作记录
	//	systemRouter.InitSysDictionaryDetailRouter(PrivateGroup) // 字典详情管理
	//}

	//InstallPlugin(PublicGroup, PrivateGroup) // 安装插件

	global.GVA_LOG.Info("router register success")
	return Router
}
