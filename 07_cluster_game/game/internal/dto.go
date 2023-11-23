package internal

type PlayerInfoReq struct {
	ID int64
}

type PlayerInfoResp struct {
	ID    int64
	Name  string
	Level int
	Exp   int64
}

type TaskReq struct {
	ID int64
}

type PlayerUpdateInfo struct {
	Player *PlayerInfoResp
	Info   string
}

// ChatJoin 加入群组
type ChatJoin struct {
	UID     int64
	GroupID int64
}
