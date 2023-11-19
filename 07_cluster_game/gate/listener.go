package gate

import (
	"github.com/topfreegames/pitaya/v2/cluster"
	"sync"
)

type gameServerListener struct {
	lock         sync.RWMutex
	gameSvrIDMap map[string]string // 从 (数据库ID) account.gameplayer.serverID 到 (动态) cluster.Server.ID 的映射
}

const (
	ServerTypeOfGame  = "game"
	MetaKeyOfServerID = "server_id"
)

var _ cluster.SDListener = (*gameServerListener)(nil)

func newGameListener() *gameServerListener {
	return &gameServerListener{
		lock:         sync.RWMutex{},
		gameSvrIDMap: map[string]string{},
	}
}

func (l *gameServerListener) AddServer(svr *cluster.Server) {
	if svr.Type == ServerTypeOfGame {
		if svrID, ok := svr.Metadata[MetaKeyOfServerID]; ok {
			l.lock.Lock()
			defer l.lock.Unlock()
			// why add to map?
			// 未正确关闭服务器的情况下，可能会加入多个 svrID 相同的 svr 实例，
			// 需等待不可用的server从etcd中移除，客户端的请求才能路由到正确的game server
			l.gameSvrIDMap[svrID] = svr.ID
		}
	}
}

func (l *gameServerListener) RemoveServer(svr *cluster.Server) {
	if svr.Type == ServerTypeOfGame {
		if svrID, ok := svr.Metadata[MetaKeyOfServerID]; ok {
			l.lock.Lock()
			defer l.lock.Unlock()
			// 如果值不相同，则不能移除
			if l.gameSvrIDMap[svrID] == svr.ID {
				// delete
				delete(l.gameSvrIDMap, svrID)
			}
		}
	}
}

func (l *gameServerListener) GetGameServerID(svrID string) (string, bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	// get server.ID
	val, ok := l.gameSvrIDMap[svrID]
	return val, ok
}
