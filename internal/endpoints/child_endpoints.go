package endpoints

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"young-eagles/external/models"
	"young-eagles/internal/services"
)

type ChildEndpoints struct {
	PostChildEndpoint endpoint.Endpoint
	GetChildEndpoint  endpoint.Endpoint
}

func NewChildEndpoints(s services.ChildrenService) ChildEndpoints {
	return ChildEndpoints{
		PostChildEndpoint: MakePostChildEndpoint(s),
		GetChildEndpoint:  MakeGetChildEndpoint(s),
	}
}

type PostChildRequest struct {
	Body PostChildRequestBody
}

type PostChildRequestBody struct {
	Ch models.Child `json:"child"`
}

type postChildResponse struct {
	Body PostChildResponseBody
}

type PostChildResponseBody struct {
	Child models.Child `json:"child"`
}

func MakePostChildEndpoint(s services.ChildrenService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(PostChildRequest)
		if !ok {
			return nil, errors.New("cannot cast request to PostChildRequest")
		}

		child, err := s.PostChildDatum(ctx, req.Body.Ch)
		if err != nil {
			return nil, err
		}

		return postChildResponse{Body: PostChildResponseBody{Child: *child}}, nil
	}
}

type ChildGetRequest struct {
	ChildUUID uuid.UUID `json:"childUUID"`
}

type getChildResponse struct {
	Body childResponse
}

type childResponse struct {
	Child models.Child `json:"child"`
}

func MakeGetChildEndpoint(s services.ChildrenService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(ChildGetRequest)
		if !ok {
			return nil, errors.New("cannot cast request to ChildGetRequest")
		}
		c, err := s.GetChildByUUID(ctx, req.ChildUUID.String())
		if err != nil {
			return nil, err
		}
		return getChildResponse{Body: childResponse{Child: *c}}, nil
	}
}
