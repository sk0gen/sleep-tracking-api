package gapi

import (
	"context"
	"github.com/sk0gen/sleep-tracking-api/internal/pb"
	"github.com/sk0gen/sleep-tracking-api/util"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginResponse, error) {
	user, err := s.store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(req.Password, user.PasswordHash)
	if err != nil {
		return nil, err
	}

	token, err := s.jwtMaker.CreateToken(user.ID, s.config.AuthConfig.JWTTokenExpiration)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Token: token,
		User: &pb.UserResponse{
			Username:  user.Username,
			CreatedAt: timestamppb.New(user.CreatedAt),
		},
	}, nil
}
