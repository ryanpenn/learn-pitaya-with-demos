package internal

type LoginReq struct {
	Name     string
	Password string
}

type LoginResp struct {
	Code  int
	Token string
	Addr  string
}

type ServerInfo struct {
	ServerID   string
	ServerName string
	ServerKey  string // uuid
	Role       *RoleInfo
}

type RoleInfo struct {
	RoleID int64
	Name   string
}
