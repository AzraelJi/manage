package core

import (
	"fmt"
	"go.uber.org/zap"
	"time"
	"zga-client-manage/global"
	"zga-client-manage/initialize"
	"zga-client-manage/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
