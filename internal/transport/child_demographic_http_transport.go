package transport

import (
	"context"
	"encoding/json"
	"github.com/friendsofgo/errors"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"young-eagles/internal/endpoints"
)

func PostChildInformation(endpoint endpoints.ChildEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/child/addChild",
		httptransport.NewServer(
			endpoint.PostChildEndpoint,
			decodePostChildDatum,
			encodeResponse,
			options...,
		)).Methods(http.MethodPost)
}

func decodePostChildDatum(_ context.Context, r *http.Request) (interface{}, error) {
	var childDemo endpoints.PostChildRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("unable to read request body for child post")
	}

	if err := json.Unmarshal(body, &childDemo.Body); err != nil {
		return nil, errors.New("unable to unmarshal body")
	}

	if len(childDemo.Body.Ch.FirstName) == 0 || len(childDemo.Body.Ch.LastName) == 0 {
		return nil, errors.New("invalid request body")
	}

	return childDemo, nil
}

func GetChildInformation(endpoint endpoints.ChildEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/child/{childUuid}",
		httptransport.NewServer(
			endpoint.GetChildEndpoint,
			decodeGetChildDatum,
			encodeResponse,
			options...,
		)).Methods(http.MethodGet)
}

func decodeGetChildDatum(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.ChildGetRequest
	vars := mux.Vars(r)
	childUuid, err := uuid.Parse(vars["childUuid"])
	if err != nil {
		return nil, errors.New("invalid uuid")
	}
	req.ChildUUID = childUuid
	return req, nil
}
