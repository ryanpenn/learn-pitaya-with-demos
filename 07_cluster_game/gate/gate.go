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
	MetaKeyOfConnects        = "connects"
	MetaKeyOfPort            = "port"
	SessionKeyOfGameServerID = "game_server_id"
	SessionKeyOfAuth         = "auth"
)

// sessionHandler 监听客户端的连接和断开
func sessionHandler(c *config.GateConfig, app pitaya.Pitaya, pool session.SessionPool, l *gameServerListener) {
	// init metadata
	app.GetServer().Metadata[MetaKeyOfConnects] = strconv.Itoa(0)
	app.GetServer().Metadata[MetaKeyOfPort] = strconv.Itoa(c.Port)

	// session handler
	var lock sync.Mutex
	pool.OnAfterSessionBind(func(_ context.Context, _ session.Session) error {
		lock.Lock()
		defer lock.Unlock()

		app.GetServer().Metadata[MetaKeyOfConnects] = strconv.FormatInt(pool.GetSessionCount(), 10)
		return nil
	})
	pool.OnSessionClose(func(s session.Session) {
		lock.Lock()
		defer lock.Unlock()

		app.GetServer().Metadata[MetaKeyOfConnects] = strconv.FormatInt(pool.GetSessionCount(), 10)
		// 获取game服ID
		svrID := s.Get(SessionKeyOfGameServerID)
		uid := s.UID()
		logger.Log.Debug("掉线 uid=%s,area = %d ", s.UID(), svrID)
		if uid != "" && svrID != nil {
			userID, err := strconv.ParseInt(uid, 10, 64)
			if err != nil {
				return
			}
			if err != nil {
				logger.Log.Debug("掉线 uid=%s,area = %s servers 错误= %v ", s.UID(), svrID, err)
				return
			}
			if key, ok := l.GetGameServerID(fmt.Sprintf("%v", svrID)); ok {
				go func() {
					err := app.RPCTo(context.TODO(), key, "game.remote.loginout", nil, &protos.RPCMsg{
						Code: userID,
					})
					if err != nil {
						logger.Log.Debug("掉线 uid=%s,area = %s RPC错误= %v ", s.UID(), svrID, err)
					}
				}()
				return
			}
		}
	})
}

// addGameRouter 为game服添加路由
func addGameRouter(app pitaya.Pitaya, serverType string, l *gameServerListener) {
	app.AddRoute(serverType, func(ctx context.Context, route *route.Route, payload []byte, servers map[string]*cluster.Server) (*cluster.Server, error) {
		logger.Log.Debug("request router.= ", route.String())
		s := app.GetSessionFromCtx(ctx)

		// 检查是否登录
		auth := s.Get(SessionKeyOfAuth)
		if auth == nil || !auth.(bool) {
			return nil, errors.NewError(fmt.Errorf("no auth"), "403")
		}

		// 获取game服ID
		svrID := s.Get(SessionKeyOfGameServerID)
		if svrID != nil {
			if key, ok := l.GetGameServerID(fmt.Sprintf("%v", svrID)); ok {
				if svr, ok := servers[key]; ok {
					return svr, nil
				}
			}
		}

		return nil, errors.NewError(fmt.Errorf("server unavailable"), "500")
	})
}
