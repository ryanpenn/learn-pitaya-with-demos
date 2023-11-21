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
	ServerID   string
	ServerName string
	ServerKey  string // uuid
	Role       *RoleInfo
}

type RoleInfo struct {
	RoleID int64
	Name   string
}

type PlayerInfoReq struct {
	ID int64
}

type PlayerInfo struct {
	ID    int64
	Name  string
	Level int
	Exp   int64
}

type TaskReq struct {
	ID int64
}

type PlayerUpdateInfo struct {
	Player *PlayerInfo
	Info   string
}

type ChatMessage struct {
	ChatType int // 0:单聊 1:单服群里 2:跨服群聊
	From     int64
	To       int64
	Content  string
	Time     int64
}
