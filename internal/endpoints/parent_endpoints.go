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
	PostParentEndpoint endpoint.Endpoint
	GetParentEndpoint  endpoint.Endpoint
}

func MakeParentEndpoints(s services.ParentService) ParentEndpoints {
	return ParentEndpoints{
		PostParentEndpoint: MakePostParentEndpoint(s),
		GetParentEndpoint:  MakeGetParentEndpoint(s),
	}
}

type PostParentRequest struct {
	Body PostParentRequestBody
}

type PostParentRequestBody struct {
	ParentFirstName   *string `json:"parentFirstName"`
	ParentLastName    *string `json:"parentLastName"`
	ParentEmail       *string `json:"parentEmail"`
	ParentPhoneNumber string  `json:"parentPhoneNumber"`
}

type PostParentResponse struct {
	Body PostParentResponseBody
}
type PostParentResponseBody struct {
	Parent models.Parent `json:"parent"`
}

func MakePostParentEndpoint(s services.ParentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req, ok := request.(PostParentRequest)
		if !ok {
			return nil, errors.New("cannot cast request to PostParentRequest")
		}

		pM, err := s.PostParentData(ctx, *req.Body.ParentFirstName, *req.Body.ParentLastName, *req.Body.ParentEmail, req.Body.ParentPhoneNumber)
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
	Parent models.Parent `json:"parent"`
}

func MakeGetParentEndpoint(s services.ParentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(GetParentRequest)
		if !ok {
			return nil, errors.New("cannot cast request to GetParentRequest")
		}

		pM, err := s.GetParentData(ctx, req.ParentUUID)
		if err != nil {
			return nil, err
		}

		return GetParentResponse{GetParentResponseBody{Parent: *pM}}, nil
	}
}
