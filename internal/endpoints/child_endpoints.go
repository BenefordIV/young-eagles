package endpoints

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"young-eagles/external/models"
	"young-eagles/internal/services"
)

type ChildEndpoints struct {
	PostChildEndpoint endpoint.Endpoint
}

func MakeChildEndpoints(s services.ChildrenService) ChildEndpoints {
	return ChildEndpoints{
		PostChildEndpoint: MakePostChildEndpoint(s),
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
