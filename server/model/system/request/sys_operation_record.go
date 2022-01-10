package request

import (
	"zga-client-manage/model/common/request"
	"zga-client-manage/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
