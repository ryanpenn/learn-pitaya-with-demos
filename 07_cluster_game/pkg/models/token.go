package models

type AccountToken struct {
	AppID string // 应用ID
	AccID string // 账号ID
	SvrID int32  // 服务器ID
	UID   int64  // 游戏角色ID
	UUID  string // 客户端ID
}
