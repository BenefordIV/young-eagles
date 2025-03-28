package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"log"
	"young-eagles/external/models"
	"young-eagles/internal/services"
)

type PlaneEndpoints struct {
	AddPlaneEndpoint       endpoint.Endpoint
	DeletePlaneEndpoint    endpoint.Endpoint
	ReinstatePlaneEndpoint endpoint.Endpoint
}

func MakePlaneEndpoints(s services.PlanesService) PlaneEndpoints {
	return PlaneEndpoints{
		AddPlaneEndpoint:       MakeAddPlaneEndpoint(s),
		DeletePlaneEndpoint:    MakeDeletePlaneEndpoint(s),
		ReinstatePlaneEndpoint: MakeReinstatePlaneDatumEndpoint(s),
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
		p, err := s.AddPlaneDatum(ctx, req.Body.CallNumber, req.Body.PlaneModel, req.Body.PlaneMake)
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

type ReinstatePlaneDatumRequest struct {
	CallNumber string `json:"callNumber"`
}

func MakeReinstatePlaneDatumEndpoint(s services.PlanesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(ReinstatePlaneDatumRequest)
		if !ok {
			return nil, nil
		}

		log.Println("ReinstatePlaneDatumEndpoint called")
		plane, err := s.ReinstatePlaneDatum(ctx, req.CallNumber)
		if err != nil {
			return nil, err
		}

		resp := postPlaneResponse{
			Body: postPlaneResponseBody{
				Plane: *plane,
			},
		}

		return resp.Body, nil
	}
}
