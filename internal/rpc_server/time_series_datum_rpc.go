package rpc_server

import (
	c "github.com/mikaponics/mikapod-storage/pkg/rpc_client"
)

func (rpc *RPC) AddTimeSeriesDatum(request *c.TimeSeriesDatumCreateRequest, response *c.TimeSeriesDatumCreateResponse) error {
	err := rpc.Store.InsertTimeSeriesData(request.Instrument, request.Value, request.Timestamp)
	if err != nil {
		*response = *&c.TimeSeriesDatumCreateResponse{
			Status: true,
		}
	}
	return err
}

func (s *RPC) ListTimeSeriesData(request *c.TimeSeriesDatumListRequest, response *c.TimeSeriesDatumListResponse) error {
	// Call our local storage and return all the results we have saved.
	rawResults := s.Store.ListTimeSeriesData()

	// Iterate through all our results and generate our results payload to send to client.
	var results []*c.TimeSeriesDatumListItemResponse
	for _, val := range rawResults {
		item := &c.TimeSeriesDatumListItemResponse{
			Id:         val.Id,
			Instrument: val.Instrument,
			Value:      val.Value,
			Timestamp:  val.Timestamp,
		}
		results = append(results, item)
	}

	// Send the data to the client.
	*response = c.TimeSeriesDatumListResponse{ // Return through RPC.
		Results: results,
	}

	return nil
}

func (s *RPC) DeleteTimeSeriesData(request *c.TimeSeriesDatumDeleteRequest, response *c.TimeSeriesDatumDeleteResponse) error {
	// Call our local storage and delete all the records for the inputted pks.
	s.Store.DeleteTimeSeriesData(request.Pks)

	// Send success message to client.
	*response = c.TimeSeriesDatumDeleteResponse{
		Status: true,
	}

	return nil
}
