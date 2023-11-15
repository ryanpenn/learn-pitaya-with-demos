package services

import (
	"context"

	pb "learn-pitaya-with-demos/cluster_chat/protos"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
)

type Account struct {
	component.Base
	app pitaya.Pitaya
}

// TODO 暂时用这个存贮账号密码
var AccountMap map[string]string = map[string]string{
	"123123": "123123",
	"abc":    "abc",
}

// 实例化一个句柄
func NewAccount(app pitaya.Pitaya) *Account {
	return &Account{app: app}
}

// 登录参数
type ReqLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// req就用json
// proto返还就用demo里面的了
func (a *Account) Login(ctx context.Context, req *ReqLogin) (*pb.Response, error) {

	// 没找到用户
	password, ok := AccountMap[req.UserName]
	if !ok {
		return ReplayError("account not found")
	}

	// 密码错误
	if password != req.Password {
		return ReplayError("password error")
	}

	// 登录成功了的话，绑定一下我的uid
	session_ := a.app.GetSessionFromCtx(ctx)
	err := session_.Bind(ctx, req.UserName)
	if err != nil {
		return ReplayError("bind session error")
	}

	return &pb.Response{
		Code: 0,
		Msg:  "success",
	}, nil
}
