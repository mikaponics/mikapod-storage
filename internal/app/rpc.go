package app

import (
	"context"
	// "log"

	tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/mikaponics/mikapod-storage/api"
	"github.com/mikaponics/mikapod-storage/internal/storage"
)

type MikapodStorageGRPC struct{
	db *storage.MikapodDB
}

func (s *MikapodStorageGRPC) AddTimeSeriesDatum(ctx context.Context, in *pb.TimeSeriesDatumRequest) (*pb.MikapodStorageResponse, error) {
	s.db.InsertTimeSeriesData(in.Instrument, in.Value, in.Timestamp)
	return &pb.MikapodStorageResponse{
		Message: "Instrument was updated",
		Status: true,
	}, nil
}

func (s *MikapodStorageGRPC) ListTimeSeriesDatum(ctx context.Context, in *pb.ListTimeSeriesDataRequest) (*pb.ListTimeSeriesDataResponse, error) {
	// Fetch the data from our database.
	data := s.db.ListTimeSeriesData()

	// // For debugging purposes only.
	// log.Printf("data: %v", data)

    // Convert our `struct` formatted list to be of `protocol buffer`
	// formatted list which we can use in our `grpc` output.
	var list []*pb.TimeSeriesDatum
	for _, v := range data {
		// Convert from int64 to `protocol buffer timestamp` object.
		ts := &tspb.Timestamp{
		    Seconds: v.Timestamp,
		    Nanos: 0,
	    }

        // Create our `protocol buffer` single time-series datum object.
        ri := &pb.TimeSeriesDatum{
            Id:         v.Id,
            Instrument: v.Instrument,
            Value:      v.Value,
			Timestamp:  ts,
        }

		// Attach our single time-series datum object to our `protocol buffer`
		// list of time-series data.
        list = append(list, ri)
    }

	return &pb.ListTimeSeriesDataResponse{
		Data: list,
	}, nil
}
