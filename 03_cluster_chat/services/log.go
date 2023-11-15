package services

import (
	"context"
	"fmt"

	pb "learn-pitaya-with-demos/cluster_chat/protos"

	"github.com/topfreegames/pitaya/v2/component"
)

type Log struct {
	component.Base
}

func (log *Log) RecordLog(ctx context.Context, req *pb.Response) (*pb.Response, error) {
	fmt.Println("record log:", req)
	return &pb.Response{
		Code: 0,
		Msg:  "Success",
	}, nil
}
