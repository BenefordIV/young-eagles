package transport

import (
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"young-eagles/internal/endpoints"
)

func AddPlaneDatum(endpoint endpoints.PlaneEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/plane/addPlane",
		httptransport.NewServer(
			endpoint.AddPlaneEndpoint,
			decodeAddPlaneDatum,
			encodeResponse,
			options...,
		)).Methods(http.MethodPost)
}

func decodeAddPlaneDatum(_ context.Context, req *http.Request) (interface{}, error) {
	var request endpoints.PostPlaneRequest
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.New("unable to read body")
	}
	if err := json.Unmarshal(body, &request.Body); err != nil {
		return nil, errors.New("unable to unmarshal body")
	}

	if len(request.Body.CallNumber) == 0 {
		log.Println("CallNumber is a Required Field. It cannot be empty!")
		return nil, errors.New("invalid request body")
	}

	return request, nil
}

func DeletePlaneDatum(endpoint endpoints.PlaneEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/plane/{callNumber}",
		httptransport.NewServer(
			endpoint.DeletePlaneEndpoint,
			decodeDeletePlaneDatum,
			encodeResponse,
			options...,
		)).Methods(http.MethodDelete)
}

func decodeDeletePlaneDatum(_ context.Context, request *http.Request) (interface{}, error) {
	req := endpoints.DeletePlaneRequest{}

	vars := mux.Vars(request)
	if callNumber, ok := vars["callNumber"]; ok {
		req.CallNumber = callNumber
	} else {
		return nil, errors.New("callNumber is a required parameter")
	}

	return req, nil
}

func ReinstatePlaneDatum(endpoint endpoints.PlaneEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/plane/reinstate/{callNumber}",
		httptransport.NewServer(
			endpoint.ReinstatePlaneEndpoint,
			decodeReinstatePlaneDatum,
			encodeResponse,
			options...,
		)).Methods(http.MethodPatch)
}

func decodeReinstatePlaneDatum(_ context.Context, request *http.Request) (interface{}, error) {
	req := endpoints.ReinstatePlaneDatumRequest{}

	vars := mux.Vars(request)
	if callNumber, ok := vars["callNumber"]; ok {
		req.CallNumber = callNumber
	} else {
		return nil, errors.New("callNumber is a required parameter")
	}

	return req, nil
}
