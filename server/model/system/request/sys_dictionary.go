package request

import (
	"zga-client-manage/model/common/request"
	"zga-client-manage/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
