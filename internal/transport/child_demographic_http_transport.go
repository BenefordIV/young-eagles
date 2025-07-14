package transport

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"young-eagles/internal/endpoints"
)

func PostChildInformation(endpoint endpoints.ChildEndpoints, router *mux.Router) {
	options := []httptransport.ServerOption{}

	router.Handle(
		"/pilot/addChild",
		httptransport.NewServer(
			endpoint.PostChildEndpoint,
			decodePostChildDatum,
			encodeResponse,
			options...,
		)).Methods(http.MethodPost)
}

func decodePostChildDatum(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
