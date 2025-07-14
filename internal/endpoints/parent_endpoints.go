package endpoints

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"young-eagles/external/models"
	"young-eagles/internal/services"
)

type ParentEndpoints struct {
	PostParentEndpoint              endpoint.Endpoint
	GetParentEndpoint               endpoint.Endpoint
	PostParentWithChildDataEndpoint endpoint.Endpoint
}

func MakeParentEndpoints(s services.ParentService) ParentEndpoints {
	return ParentEndpoints{
		PostParentEndpoint:              MakePostParentEndpoint(s),
		GetParentEndpoint:               MakeGetParentEndpoint(s),
		PostParentWithChildDataEndpoint: MakePostParentWithKidDataEndpoint(s),
	}
}

type PostParentRequest struct {
	Body PostParentRequestBody
}

type PostParentRequestBody struct {
	Par models.Parent
}

type PostParentResponse struct {
	Body PostParentResponseBody
}
type PostParentResponseBody struct {
	Parent models.Parent `json:"parent"`
}

func MakePostParentEndpoint(s services.ParentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req, ok := request.(PostParentRequestBody)
		if !ok {
			return nil, errors.New("cannot cast request to PostParentRequest")
		}

		pM, err := s.PostParentData(ctx, req.Par)
		if err != nil {
			return nil, err
		}

		return PostParentResponse{PostParentResponseBody{Parent: *pM}}, nil

	}
}

type GetParentRequest struct {
	ParentUUID uuid.UUID `json:"parentUUID"`
}

type GetParentResponse struct {
	Body GetParentResponseBody
}

type GetParentResponseBody struct {
	Parent   models.Parent  `json:"parent"`
	Children []models.Child `json:"children"`
}

func MakeGetParentEndpoint(s services.ParentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(GetParentRequest)
		if !ok {
			return nil, errors.New("cannot cast request to GetParentRequest")
		}

		pM, cM, err := s.GetParentData(ctx, req.ParentUUID)
		if err != nil {
			return nil, err
		}

		return GetParentResponse{GetParentResponseBody{Parent: *pM, Children: cM}}, nil
	}
}

type PostParentWithChildDataRequest struct {
	Parent   models.Parent  `json:"parent"`
	Children []models.Child `json:"children"`
}

type postParentWithChildDataResponse struct {
	Body postParentWithChildDataResponseBody `json:"body"`
}

type postParentWithChildDataResponseBody struct {
	Parent models.ParentWithChildren `json:"parent"`
}

func MakePostParentWithKidDataEndpoint(s services.ParentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(PostParentWithChildDataRequest)
		if !ok {
			return nil, errors.New("cannot cast request to PostParentWithChildDataRequest")
		}

		pM, err := s.PostParentWithChildData(ctx, req.Parent, req.Children)
		if err != nil {
			return nil, err
		}

		return postParentWithChildDataResponse{postParentWithChildDataResponseBody{Parent: *pM}}, nil
	}
}
