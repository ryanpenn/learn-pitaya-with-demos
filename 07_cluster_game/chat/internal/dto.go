package internal

// ChatMessage 聊天消息
type ChatMessage struct {
	ChatType int // 0:单聊 1:群聊
	From     int64
	To       int64
	Content  string
	Time     int64
}

// RoomInfo 群聊信息
type RoomInfo struct {
	ID       int64
	RoomType int // 0:单服群 1:联盟群 2:活动群(跨服)
	Name     string
}

// ReqChatHistory 请求聊天记录
type ReqChatHistory struct {
}

// ChatHistory 聊天记录
type ChatHistory struct {
}

// ChatJoin 加入群组
type ChatJoin struct {
	UID     int64
	GroupID int64
}
