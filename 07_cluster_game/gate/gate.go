package gate

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/cluster"
	"github.com/topfreegames/pitaya/v2/errors"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/route"
	"github.com/topfreegames/pitaya/v2/session"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
	"learn-pitaya-with-demos/cluster_game/protos"
	"strconv"
	"sync"
)

const (
	MetaKeyOfConnects = "conn"
	MetaKeyOfPort     = "port"
)

// sessionHandler 监听客户端的连接和断开
func sessionHandler(c *config.GateConfig, app pitaya.Pitaya, pool session.SessionPool) {
	// gate服连接数 和 端口
	app.GetServer().Metadata[MetaKeyOfConnects] = strconv.Itoa(0)
	app.GetServer().Metadata[MetaKeyOfPort] = strconv.Itoa(c.Port)

	// 在这里校验token
	pool.AddHandshakeValidator("game", func(data *session.HandshakeData) error {
		// 登录的game服
		if _, exist := data.User["key"]; !exist {
			return errors.NewError(fmt.Errorf("no target server"), "404")
		}

		// token
		if tk, ok := data.User["token"]; ok {
			err := app.RPC(context.Background(), "login.remote.auth", &protos.RPCEmpty{}, &protos.RPCMsg{
				Code:    0,
				Content: tk.(string),
			})
			if err != nil {
				fmt.Println("RPC err", err)
			}
			return err
		}

		return errors.NewError(fmt.Errorf("no auth"), "401")
	})

	// session handler
	var lock sync.Mutex
	pool.OnAfterSessionBind(func(_ context.Context, s session.Session) error {
		lock.Lock()
		defer lock.Unlock()

		app.GetServer().Metadata[MetaKeyOfConnects] = strconv.FormatInt(pool.GetSessionCount(), 10)
		fmt.Println("conn:", app.GetServer().Metadata[MetaKeyOfConnects])
		return nil
	})
	pool.OnSessionClose(func(s session.Session) {
		lock.Lock()
		defer lock.Unlock()

		app.GetServer().Metadata[MetaKeyOfConnects] = strconv.FormatInt(pool.GetSessionCount(), 10)
		fmt.Println("conn:", app.GetServer().Metadata[MetaKeyOfConnects])
		// 处理掉线
		handleOffline(app, s.GetHandshakeData().User["key"].(string), s.UID())
	})
}

func handleOffline(app pitaya.Pitaya, key, uid string) {
	fmt.Println("rpc", key, "game.remote.offline", uid)
	err := app.RPCTo(context.TODO(), key, "game.remote.offline", &protos.RPCEmpty{}, &protos.RPCMsg{
		Code:    0,
		Content: uid,
	})
	if err != nil {
		logger.Log.Debug("掉线 uid=%s, key=%s, RPC错误=%v ", uid, key, err)
	}
}

// addGameRouter 为game服添加路由
func addGameRouter(app pitaya.Pitaya, serverType string) {
	app.AddRoute(serverType, func(ctx context.Context, route *route.Route, payload []byte, servers map[string]*cluster.Server) (*cluster.Server, error) {
		logger.Log.Debug("request router.= ", route.String())

		s := app.GetSessionFromCtx(ctx)
		key := s.GetHandshakeData().User["key"].(string)
		if svr, ok := servers[key]; ok {
			return svr, nil
		}

		return nil, errors.NewError(fmt.Errorf("server unavailable"), "500")
	})
}
