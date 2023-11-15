package services

import (
	"errors"

	pb "learn-pitaya-with-demos/cluster_chat/protos"

	"github.com/topfreegames/pitaya/v2"
)

func ReplayError(msg string) (*pb.Response, error) {
	return &pb.Response{
		Code: 1,
		Msg:  msg,
	}, pitaya.Error(errors.New(msg), "1")
}
