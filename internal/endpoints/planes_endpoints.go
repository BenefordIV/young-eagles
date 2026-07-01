package endpoints

import (
	"context"
	"log"
	"young-eagles/external/models"
	"young-eagles/internal/services"

	"github.com/go-kit/kit/endpoint"
)

type PlaneEndpoints struct {
	AddPlaneEndpoint    endpoint.Endpoint
	DeletePlaneEndpoint endpoint.Endpoint
}

func MakePlaneEndpoints(s services.PlanesService) PlaneEndpoints {
	return PlaneEndpoints{
		AddPlaneEndpoint:    MakeAddPlaneEndpoint(s),
		DeletePlaneEndpoint: MakeDeletePlaneEndpoint(s),
	}
}

type PostPlaneRequest struct {
	Body PostPlaneRequestBody
}

type PostPlaneRequestBody struct {
	CallNumber string `json:"callNumber"`
	PlaneModel string `json:"planeModel"`
	PlaneMake  string `json:"planeMake"`
}

type postPlaneResponse struct {
	Body postPlaneResponseBody `json:"body"`
}

type postPlaneResponseBody struct {
	Plane models.Plane
}

func MakeAddPlaneEndpoint(s services.PlanesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(PostPlaneRequest)
		if !ok {
			return nil, nil
		}

		log.Println("AddPlaneEndpoint called")
		plane := models.Plane{
			CallNumber: req.Body.CallNumber,
			PlaneModel: req.Body.PlaneModel,
			PlaneMake:  req.Body.PlaneMake,
		}
		p, err := s.AddPlaneDatum(ctx, plane)
		if err != nil {
			return nil, err
		}

		resp := postPlaneResponse{
			Body: postPlaneResponseBody{
				Plane: *p,
			},
		}

		return resp.Body, nil
	}
}

type DeletePlaneRequest struct {
	CallNumber string `json:"callNumber"`
}

func MakeDeletePlaneEndpoint(s services.PlanesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(DeletePlaneRequest)
		if !ok {
			return nil, nil
		}

		log.Println("DeletePlaneEndpoint called")
		err := s.DeletePlaneDatum(ctx, req.CallNumber)
		if err != nil {
			return nil, err
		}

		return EmptyResponse{}, nil // No content response for successful deletion
	}
}
