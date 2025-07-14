package transport

import (
	"context"
	"encoding/json"
	"errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"young-eagles/internal/endpoints"
)

func PostParentInformation(endpoint endpoints.ParentEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/parent/addParent",
		httptransport.NewServer(
			endpoint.PostParentEndpoint,
			decodePostParentDatum,
			encodeResponse,
			options...,
		)).Methods(http.MethodPost)
}

func decodePostParentDatum(ctx context.Context, req *http.Request) (interface{}, error) {
	var request endpoints.PostParentRequest

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, errors.New("unable to read body")
	}
	if err := json.Unmarshal(body, &request.Body.Par); err != nil {
		return nil, errors.New("unable to unmarshal body")
	}

	if request.Body.Par.FirstName == nil || request.Body.Par.LastName == nil {
		return nil, errors.New("invalid request body: ParentFirstName and ParentLastName must not be empty")
	}

	if request.Body.Par.Email == nil {
		return nil, errors.New("invalid request body: ParentEmail must not be empty")
	}

	return req, nil
}

func GetParentDatum(endpoint endpoints.ParentEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle("/parent/getParent/%s/parentId",
		httptransport.NewServer(
			endpoint.GetParentEndpoint,
			decodeGetParentDatum,
			encodeResponse,
			options...,
		)).Methods(http.MethodGet)
}

func decodeGetParentDatum(ctx context.Context, request *http.Request) (interface{}, error) {
	vars := mux.Vars(request)
	pId := vars["parentId"]
	pUUID, err := uuid.Parse(pId)
	if err != nil {
		return nil, errors.New("invalid parent UUID format")
	}

	return endpoints.GetParentRequest{ParentUUID: pUUID}, nil
}

func PostParentWithChildrenData(endpoint endpoints.ParentEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/parent/addParentWithChildren",
		httptransport.NewServer(
			endpoint.PostParentWithChildDataEndpoint,
			decodePostParentWithChildData,
			encodeResponse,
			options...,
		)).Methods(http.MethodPost)
}

func decodePostParentWithChildData(ctx context.Context, request *http.Request) (interface{}, error) {
	var req endpoints.PostParentWithChildDataRequest

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, errors.New("unable to read body")
	}
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, errors.New("unable to unmarshal body")
	}

	if req.Parent.FirstName == nil || req.Parent.LastName == nil {
		return nil, errors.New("invalid request body: ParentFirstName and ParentLastName must not be empty")
	}

	if req.Parent.Email == nil {
		return nil, errors.New("invalid request body: ParentEmail must not be empty")
	}

	return req, nil
}
