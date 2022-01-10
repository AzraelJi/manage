package v1

import (
	"zga-client-manage/api/v1/business"
	"zga-client-manage/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	AutoCodeApiGroup business.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
