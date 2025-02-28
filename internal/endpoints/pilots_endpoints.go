package endpoints

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"young-eagles/internal/services"
)

type PilotEndpoints struct {
	PostPilotEndpoint endpoint.Endpoint
}

func MakePilotEndpoints(s services.PilotService) PilotEndpoints {
	return PilotEndpoints{
		PostPilotEndpoint: MakePostPilotEndpoint(s),
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

		p, err := s.PostPilotService(ctx, req.Body.PilotFirstName, req.Body.PilotLastName, req.Body.PilotEmail, req.Body.EaaChapterNumber)
		if err != nil {
			return nil, err
		}

		resp := postPilotResponse{
			Body: postPilotResponseBody{
				Uuid: p.PilotUuid.String(),
			},
		}

		return resp, nil
	}
}
