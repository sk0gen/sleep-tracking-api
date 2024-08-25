package gapi

import (
	db "github.com/sk0gen/sleep-tracking-api/internal/database/sqlc"
	"github.com/sk0gen/sleep-tracking-api/internal/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapSleepLogs(sleepLogs []db.SleepLog) []*pb.UserSleepLog {
	var logs = make([]*pb.UserSleepLog, 0, len(sleepLogs))
	for _, v := range sleepLogs {
		logs = append(logs, &pb.UserSleepLog{
			Id:        v.ID.String(),
			StartTime: timestamppb.New(v.StartTime),
			EndTime:   timestamppb.New(v.EndTime),
			Quality:   v.Quality,
			CreatedAt: timestamppb.New(v.CreatedAt),
		})
	}

	return logs
}
