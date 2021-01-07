package rpc_client

import (
	// "time"
	// "errors"
	// "database/sql"
	"log"
	// "net/rpc"
)

// Request structure to send to the RPC server.

type TimeSeriesDatumCreateRequest struct {
    Instrument int32 `json:"instrument"`
    Value float32 `json:"value"`
    Timestamp int64 `json:"timestamp"`
}

type TimeSeriesDatumCreateResponse struct {
    Status bool `json:"status"`
}

type TimeSeriesDatumListRequest struct {}

type TimeSeriesDatumListItemResponse struct {
	Id int64
    Instrument int32 `json:"instrument"`
    Value float32 `json:"value"`
    Timestamp int64 `json:"timestamp"`
}

type TimeSeriesDatumListResponse struct {
    Results []*TimeSeriesDatumListItemResponse `json:"results"`
}

type TimeSeriesDatumDeleteRequest struct {
	Pks []int64
}

type TimeSeriesDatumDeleteResponse struct {
    Status bool `json:"status"`
}

// RPC Calls

func (s *MikapodStorageService) AddTimeSeriesDatum(datum *TimeSeriesDatumCreateRequest) (bool, error) {
	var response TimeSeriesDatumCreateResponse
	rpcErr := s.Client.Call("RPC.AddTimeSeriesDatum", *datum, &response)
	if rpcErr != nil {
		log.Println("ERROR | MikapodStorageService | AddTimeSeriesDatum | rpcErr:", rpcErr)
		return false, rpcErr
	}
	return response.Status, nil
}

func (s *MikapodStorageService) ListTimeSeriesData() (*TimeSeriesDatumListResponse, error) {
	var response TimeSeriesDatumListResponse
	request := TimeSeriesDatumListRequest{}
	err := s.Client.Call("RPC.ListTimeSeriesData", request, &response)
	if err != nil {
		log.Println("ERROR | MikapodStorageService | ListTimeSeriesData | err", err)
		return nil, err
	}
	return &response, nil
}

func (s *MikapodStorageService) DeleteTimeSeriesData(pks []int64) (*TimeSeriesDatumDeleteResponse, error) {
	var response TimeSeriesDatumDeleteResponse
	request := TimeSeriesDatumDeleteRequest{
		Pks: pks,
	}
	err := s.Client.Call("RPC.DeleteTimeSeriesData", request, &response)
	if err != nil {
		log.Println("ERROR | MikapodStorageService | DeleteTimeSeriesData | err", err)
		return nil, err
	}
	return &response, nil
}
