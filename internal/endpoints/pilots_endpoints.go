package endpoints

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"young-eagles/external/models"
	"young-eagles/internal/services"
)

type PilotEndpoints struct {
	PostPilotEndpoint      endpoint.Endpoint
	GetPilotDataEndpoint   endpoint.Endpoint
	PatchPilotDataEndpoint endpoint.Endpoint
}

func MakePilotEndpoints(s services.PilotService) PilotEndpoints {
	return PilotEndpoints{
		PostPilotEndpoint:      MakePostPilotEndpoint(s),
		GetPilotDataEndpoint:   MakeGetPilotDataEndpoint(s),
		PatchPilotDataEndpoint: MakePatchPilotDataEndpoint(s),
	}
}

type postPilotResponse struct {
	Body postPilotResponseBody `json:"body"`
}

type postPilotResponseBody struct {
	Uuid string `json:"uuid"`
}

type PilotPostRequest struct {
	Body PilotPostRequestBody
}

type PilotPostRequestBody struct {
	PilotFirstName   string `json:"pilotFirstName"`
	PilotLastName    string `json:"pilotLastName"`
	EaaChapterNumber int    `json:"eaaChapterNumber"`
	PilotEmail       string `json:"pilotEmail"`
}

func MakePostPilotEndpoint(s services.PilotService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(PilotPostRequest)
		if !ok {
			return nil, errors.New("cannot cast request to models.PilotPostRequest")
		}

		p, err := s.PostPilotData(ctx, req.Body.PilotFirstName, req.Body.PilotLastName, req.Body.PilotEmail, req.Body.EaaChapterNumber)
		if err != nil {
			return nil, err
		}

		resp := postPilotResponse{
			Body: postPilotResponseBody{
				Uuid: p.PilotUuid.String(),
			},
		}

		return resp.Body, nil
	}
}

type PilotGetRequest struct {
	PilotUUID uuid.UUID `json:"pilotUUID"`
}

type getPilotResponse struct {
	Body pilotResponse
}

type pilotResponse struct {
	Pilot models.Pilot `json:"pilot"`
}

func MakeGetPilotDataEndpoint(s services.PilotService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(PilotGetRequest)
		if !ok {
			return nil, errors.New("cannot cast request to models.PilotGetRequest")
		}

		p, err := s.GetPilotData(ctx, req.PilotUUID.String())
		if err != nil {
			return nil, err
		}

		resp := getPilotResponse{
			Body: pilotResponse{
				Pilot: *p,
			},
		}
		return resp.Body, nil
	}
}

type PatchPilotRequest struct {
	PilotUUID uuid.UUID                    `json:"pilotUuid"`
	Body      models.PatchPilotBodyRequest `json:"body"`
}

func MakePatchPilotDataEndpoint(s services.PilotService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(PatchPilotRequest)
		if !ok {
			return nil, errors.New("cannot cast request to models.PatchPilotRequest")
		}

		p, err := s.PatchUpdatePilotData(ctx, req.PilotUUID.String(), req.Body)
		if err != nil {
			return nil, err
		}

		resp := getPilotResponse{
			Body: pilotResponse{
				Pilot: *p,
			},
		}

		return resp, nil
	}
}
