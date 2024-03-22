package service

import (
	"context"

	pb "realworld/api/helloworld/v1"
)

type RealWorldService struct {
	pb.UnimplementedRealWorldServer
}

func NewRealWorldService() *RealWorldService {
	return &RealWorldService{}
}

func (s *RealWorldService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *pb.GetCurrentUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) UnFollowUser(ctx context.Context, req *pb.UnFollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}
func (s *RealWorldService) FeedArticles(ctx context.Context, req *pb.FeedArticlesRequest) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}
func (s *RealWorldService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.SingleCommentReply, error) {
	return &pb.SingleCommentReply{}, nil
}
func (s *RealWorldService) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.MultipleCommentsReply, error) {
	return &pb.MultipleCommentsReply{}, nil
}
func (s *RealWorldService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.SingleCommentReply, error) {
	return &pb.SingleCommentReply{}, nil
}
func (s *RealWorldService) FavoriteArticle(ctx context.Context, req *pb.FavoriteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}
func (s *RealWorldService) UnFavoriteArticle(ctx context.Context, req *pb.UnFavoriteArticleRequest) (*pb.SingleCommentReply, error) {
	return &pb.SingleCommentReply{}, nil
}
func (s *RealWorldService) GetTags(ctx context.Context, req *pb.GetTagsRequest) (*pb.ListOfTags, error) {
	return &pb.ListOfTags{}, nil
}
