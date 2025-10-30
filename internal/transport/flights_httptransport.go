package transport

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"young-eagles/internal/endpoints"

	"github.com/friendsofgo/errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func PostFlightData(endpoint endpoints.FlightEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/flights/postFlight",
		httptransport.NewServer(
			endpoint.PostFlightEndpoint,
			decodePostFlightRequest,
			encodeResponse,
			options...,
		)).Methods(http.MethodPost)
}

func decodePostFlightRequest(_ context.Context, req *http.Request) (interface{}, error) {
	var r endpoints.PostFlightRequest
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &r.Body); err != nil {
		return nil, err
	}

	if &r.Body.Child == nil || &r.Body.Pilot == nil {
		return nil, errors.New("child or pilot cannot be null")
	}

	return r.Body, nil
}

func PatchFlightCompleted(endpoint endpoints.FlightEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/flights/{uuid}",
		httptransport.NewServer(
			endpoint.PatchFlightCompletedEndpoint,
			decodePatchFlightCompletedRequest,
			encodeResponse,
			options...,
		)).Methods(http.MethodPatch)
}

func decodePatchFlightCompletedRequest(_ context.Context, req *http.Request) (interface{}, error) {
	var r endpoints.PatchFlightCompletedRequest
	vars := mux.Vars(req)
	id, err := uuid.Parse(vars["uuid"])
	if err != nil {
		return nil, errors.New("invalid uuid")
	}
	r.FlightUUID = id
	return r, nil
}
