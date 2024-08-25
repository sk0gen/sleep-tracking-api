package gapi

import (
	"context"
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func pageDefaults(req *pb.GetUserSleepLogsRequest) {
	if req.PageNumber == 0 {
		req.PageNumber = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
}

func (s *Server) GetUserSleepLogs(ctx context.Context, req *pb.GetUserSleepLogsRequest) (*pb.GetUserSleepLogsResponse, error) {
	user, err := s.authorizeUser(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized: %s", err)
	}
	pageDefaults(req)

	res, err := s.store.GetSleepLogsByUserID(ctx, db.GetSleepLogsByUserIDParams{
		UserID: user.UserID,
		Limit:  req.PageSize,
		Offset: (req.PageNumber - 1) * req.PageSize,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get sleep logs: %s", err)
	}

	return &pb.GetUserSleepLogsResponse{
		SleepLogs: mapSleepLogs(res),
	}, nil
}
