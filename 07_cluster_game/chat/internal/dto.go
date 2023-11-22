package internal

type ChatMessage struct {
	ChatType int // 0:单聊 1:群聊
	From     int64
	To       int64
	Content  string
	Time     int64
}

type RoomInfo struct {
	ID   int64
	Name string
}
