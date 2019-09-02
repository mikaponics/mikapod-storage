package rpc

import (
	"context"

	pb "github.com/mikaponics/mikapod-storage/api"
	"github.com/mikaponics/mikapod-storage/internal/storage"
)

type MikapodStorageGRPC struct{}

func (s *MikapodStorageGRPC) AddTimeSeriesDatum(ctx context.Context, in *pb.TimeSeriesDatumRequest) (*pb.MikapodStorageResponse, error) {
	ms := storage.InitMikapodStorage()
	ms.InsertTimeSeriesData(in.Instrument, in.Value, in.Timestamp)
	return &pb.MikapodStorageResponse{
		Message: "Instrument was updated",
		Status: true,
	}, nil
}
