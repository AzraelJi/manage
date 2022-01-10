package request

import (
	"zga-client-manage/model/common/request"
	"zga-client-manage/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
