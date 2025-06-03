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
		"/pilot/addParent",
		httptransport.NewServer(
			endpoint.PostParentEndpoint,
			decodePostParentDatum,
			encodeResponse,
			options...,
		)).Methods(http.MethodPost)
}

func decodePostParentDatum(ctx context.Context, request *http.Request) (interface{}, error) {
	var req endpoints.PostParentRequest

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, errors.New("unable to read body")
	}
	if err := json.Unmarshal(body, &request.Body); err != nil {
		return nil, errors.New("unable to unmarshal body")
	}

	if req.Body.ParentFirstName == nil || req.Body.ParentLastName == nil {
		return nil, errors.New("invalid request body: ParentFirstName and ParentLastName must not be empty")
	}

	if req.Body.ParentEmail == nil {
		return nil, errors.New("invalid request body: ParentEmail must not be empty")
	}

	return req, nil
}

func GetParentDatum(endpoint endpoints.ParentEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle("/pilot/getParent/%s/parentId",
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
