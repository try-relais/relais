package avp

import (
	"context"
	"errors"

	"github.com/Derek-X-Wang/relais-avp/pkg/log"
	"github.com/Derek-X-Wang/relais-avp/pkg/process"

	pb "github.com/Derek-X-Wang/relais-avp/pkg/proto/avp"
)

func (s *server) StartProcess(ctx context.Context, in *pb.StartProcessRequest) (*pb.StartProcessReply, error) {
	log.Infof("process einfo=%v", in.Element)
	pipeline := process.GetPipeline(in.Element.Mid)
	if pipeline == nil {
		return nil, errors.New("process: pipeline not found")
	}
	pipeline.AddElement(in.Element)
	return &pb.StartProcessReply{}, nil
}

func (s *server) StopProcess(ctx context.Context, in *pb.StopProcessRequest) (*pb.StopProcessReply, error) {
	log.Infof("publish unprocess=%v", in)
	return &pb.StopProcessReply{}, nil
}
