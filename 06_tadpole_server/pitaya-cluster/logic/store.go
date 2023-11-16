package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/topfreegames/pitaya/v2/component"

	"pitaya_tadpole/logic/protocol"
	"pitaya_tadpole/protos"
)

// Store struct
type Store struct {
	component.Base
	posLock   sync.Mutex
	positions map[int]*protocol.UpdateMessage
}

// NewStore ctor
func NewStore() *Store {
	return &Store{
		posLock:   sync.Mutex{},
		positions: make(map[int]*protocol.UpdateMessage),
	}
}

// save data
func (s *Store) Save(ctx context.Context, req *protos.Arg) (*protos.Response, error) {
	fmt.Println("Store --------> Save:", req.Msg)
	return &protos.Response{
		Code: 200,
		Msg:  "Success",
	}, nil
}

func (s *Store) GetPos(ctx context.Context, req *protos.Arg) (*protos.Response, error) {
	uid, _ := strconv.Atoi(req.Msg)

	s.posLock.Lock()
	defer s.posLock.Unlock()

	fmt.Println("Store --------> GetPos:", req.Msg)

	if pos, ok := s.positions[uid]; ok {
		rawMsg, _ := json.Marshal(pos)
		return &protos.Response{
			Code: 200,
			Msg:  string(rawMsg),
		}, nil
	}

	return &protos.Response{
		Code: 0,
		Msg:  "OK",
	}, nil
}

func (s *Store) UpdatePos(ctx context.Context, req *protos.Arg) (*protos.Response, error) {
	pos, err := parseRequest(req)
	if err != nil {
		return responseError("UpdatePos", err)
	}

	s.posLock.Lock()
	defer s.posLock.Unlock()

	s.positions[pos.ID] = pos

	return &protos.Response{
		Code: 200,
		Msg:  req.Msg,
	}, nil
}

func (s *Store) RemovePos(ctx context.Context, req *protos.Arg) (*protos.Response, error) {
	posID, err := strconv.Atoi(req.Msg)
	if err != nil {
		return responseError("RemovePos", err)
	}

	s.posLock.Lock()
	defer s.posLock.Unlock()

	delete(s.positions, posID)

	return &protos.Response{
		Code: 200,
		Msg:  req.Msg,
	}, nil
}

func parseRequest(req *protos.Arg) (*protocol.UpdateMessage, error) {
	rawData := []byte(req.Msg)
	var pos protocol.UpdateMessage
	if err := json.Unmarshal(rawData, &pos); err != nil {
		return nil, err
	}

	return &pos, nil
}

func responseError(name string, err error) (*protos.Response, error) {
	msg := fmt.Sprintf("%s err: %v", name, err)
	fmt.Println(msg)
	return &protos.Response{
		Code: -1,
		Msg:  msg,
	}, err
}
