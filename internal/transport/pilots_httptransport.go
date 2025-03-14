package transport

import (
	"context"
	"encoding/json"
	"github.com/friendsofgo/errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"young-eagles/internal/endpoints"
)

func PostPilotData(endpoint endpoints.PilotEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/pilot/addPilot",
		httptransport.NewServer(
			endpoint.PostPilotEndpoint,
			decodePostPilotData,
			encodeResponse,
			options...,
		)).Methods(http.MethodPost)
}

func decodePostPilotData(ctx context.Context, req *http.Request) (interface{}, error) {
	var request endpoints.PilotPostRequest
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.New("unable to read body")
	}
	if err := json.Unmarshal(body, &request.Body); err != nil {
		return nil, errors.New("unable to unmarshal body")
	}

	if len(request.Body.PilotFirstName) == 0 || len(request.Body.PilotLastName) == 0 {
		log.Println("PilotAddPilotFirstName and PilotLastName must not be empty")
		return nil, errors.New("invalid request body")
	}

	return request, nil
}

func GetPilotData(endpoint endpoints.PilotEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/pilot/{pilotUuid}",
		httptransport.NewServer(
			endpoint.GetPilotDataEndpoint,
			decodeGetPilotData,
			encodeResponse,
			options...)).Methods(http.MethodGet)
}

func decodeGetPilotData(ctx context.Context, request *http.Request) (interface{}, error) {
	var req endpoints.PilotGetRequest

	vars := mux.Vars(request)

	uuid, err := uuid.Parse(vars["pilotUuid"])
	if err != nil {
		return nil, errors.New("invalid uuid")
	}

	req.PilotUUID = uuid

	return req, nil

}

func PatchPilotData(endpoint endpoints.PilotEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/pilot/{pilotUuid}",
		httptransport.NewServer(
			endpoint.PatchPilotDataEndpoint,
			decodePatchPilotData,
			encodeResponse,
			options...)).Methods(http.MethodPatch)
}

func decodePatchPilotData(ctx context.Context, request *http.Request) (interface{}, error) {
	var req endpoints.PatchPilotRequest

	vars := mux.Vars(request)
	pilotUuid, err := uuid.Parse(vars["pilotUuid"])
	if err != nil {
		return nil, errors.New("invalid uuid")
	}
	req.PilotUUID = pilotUuid

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, errors.New("unable to read body")
	}
	if err := json.Unmarshal(body, &req.Body); err != nil {
		return nil, errors.New("unable to unmarshal body")
	}

	return req, nil
}
