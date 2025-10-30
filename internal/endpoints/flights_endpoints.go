package endpoints

import (
	"context"
	"errors"
	"young-eagles/external/models"
	"young-eagles/internal/services"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
)

type FlightEndpoints struct {
	PostFlightEndpoint           endpoint.Endpoint
	PatchFlightCompletedEndpoint endpoint.Endpoint
}

func MakeFlightEndpoints(s services.FlightService) FlightEndpoints {
	return FlightEndpoints{
		PostFlightEndpoint:           MakePostFlightEndpoint(s),
		PatchFlightCompletedEndpoint: MakePatchFlightCompletedEndpoint(s),
	}
}

type PostFlightRequest struct {
	Body PostFlightRequestBody
}

type PostFlightRequestBody struct {
	Pilot models.Pilot `json:"pilot"`
	Child models.Child `json:"child"`
	Plane models.Plane `json:"plane"`
}

type postFlightResponse struct {
	Body PostFlightResponseBody
}
type PostFlightResponseBody struct {
	Flight models.Flight
}

func MakePostFlightEndpoint(s services.FlightService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(PostFlightRequest)
		if !ok {
			return nil, errors.New("cannot cast request to PostFlightRequest")
		}

		flight, err := s.PostFlight(ctx, req.Body.Pilot, req.Body.Child, req.Body.Plane)

		if err != nil {
			return nil, err
		}

		return postFlightResponse{PostFlightResponseBody{Flight: *flight}}, nil
	}
}

type PatchFlightCompletedRequest struct {
	FlightUUID uuid.UUID `json:"flightUUID"`
}

type patchFLightCompletedResponse struct {
	Body emptyResponse
}

func MakePatchFlightCompletedEndpoint(s services.FlightService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(PatchFlightCompletedRequest)
		if !ok {
			return nil, errors.New("cannot cast request to PatchFlightCompletedRequest")
		}

		err := s.PatchFlightFinished(ctx, req.FlightUUID)
		if err != nil {
			return nil, err
		}

		return patchFLightCompletedResponse{Body: emptyResponse{}}, nil
	}
}
