package service

import (
	"context"

	pb "realworld/api/helloworld/v1"
	"realworld/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type RealWorldService struct {
	pb.UnimplementedRealWorldServer
	uc  *biz.UserUsecase
	sc  *biz.SocialUsecase
	log *log.Helper
}

func NewRealWorldService(rwu *biz.RealWorldUsecase) *RealWorldService {
	return &RealWorldService{}
}

func (s *RealWorldService) UnFollowUser(ctx context.Context, req *pb.UnFollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
