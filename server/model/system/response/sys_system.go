package response

import "zga-client-manage/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
