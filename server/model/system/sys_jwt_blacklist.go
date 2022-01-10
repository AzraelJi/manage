package system

import (
	"zga-client-manage/global"
)

type JwtBlacklist struct {
	global.ZgaModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
