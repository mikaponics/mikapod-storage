package rpc_server

import (
	"time"
)

type TimeSeriesDatum struct {
    Id string `json:"id,omitempty"`
    Instrument int32 `json:"instrument"`
    Value float32 `json:"value"`
    Timestamp time.Time `json:"timestamp"`
}

func (s *RPC) AddTimeSeriesDatum(request *TimeSeriesDatum, response *TimeSeriesDatum) (error) {
	return nil
}

func (s *RPC) ListTimeSeriesData() (error) {
	return nil
}

func (s *RPC) DeleteTimeSeriesData() (error) {
	return nil
}
