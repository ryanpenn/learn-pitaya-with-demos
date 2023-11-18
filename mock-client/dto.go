package main

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
	ServerID   int
	ServerName string
	Role       *RoleInfo
}

type RoleInfo struct {
	RoleID int64
	Name   string
}
