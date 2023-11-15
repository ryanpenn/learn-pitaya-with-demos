package services

import (
	"context"
	"fmt"

	pb "learn-pitaya-with-demos/cluster_chat/protos"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
)

type Room struct {
	component.Base
	app pitaya.Pitaya
}

var RoomGroupName string = "RoomGroupName"

// 句柄
func NewRoom(app pitaya.Pitaya) *Room {

	app.GroupCreate(context.Background(), RoomGroupName)

	return &Room{
		app: app,
	}
}

// 加入房间
func (room *Room) Join(ctx context.Context, req []byte) (*pb.Response, error) {

	session_ := room.app.GetSessionFromCtx(ctx)
	uid := session_.UID()

	if uid == "" {
		return ReplayError("not login")
	}

	if have, _ := room.app.GroupContainsMember(ctx, RoomGroupName, uid); have {
		return ReplayError("i in room , uid:" + uid)
	}

	room.app.GroupAddMember(ctx, RoomGroupName, uid)
	room.app.GroupBroadcast(ctx, "game", RoomGroupName, "joinPush", pb.Response{
		Msg: fmt.Sprintf("------------------\n--user: %sjoin\n------------------\n", uid),
	})

	// 为了方便，请求和返还，都用response...
	replay := pb.Response{}
	room.app.ReliableRPC("log.log.recordlog", nil, &replay, &pb.Response{Msg: "uid:" + uid + ",join room"})
	// 重要‼️
	// 被ReliableRPC调用的接口必须是幂等的
	// ReliableRPC没有返回值，只有在成功时会返回作业ID（jid）

	return &pb.Response{
		Code: 0,
		Msg:  "Join Success",
	}, nil
}

// 聊天参数
type ReqMessage struct {
	Msg string `json:"msg"`
}

// 发送消息
func (room *Room) Message(ctx context.Context, req *ReqMessage) (*pb.Response, error) {

	session_ := room.app.GetSessionFromCtx(ctx)
	uid := session_.UID()

	if uid == "" {
		return ReplayError("not login")
	}

	if have, _ := room.app.GroupContainsMember(ctx, RoomGroupName, uid); !have {
		return ReplayError("i not in room")
	}

	room.app.GroupBroadcast(ctx, "game", RoomGroupName, "messagePush", pb.Response{
		Msg: fmt.Sprintf("%s say: %s\n", uid, req.Msg),
	})

	// 为了方便，请求和返还，都用response...
	replay := pb.Response{}
	room.app.ReliableRPC("log.log.recordlog", nil, &replay, &pb.Response{Msg: "uid:" + uid + ", send message :" + req.Msg})

	return &pb.Response{
		Code: 0,
		Msg:  "Send Message Success",
	}, nil
}
