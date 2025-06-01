package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"young-eagles/internal/config"
	"young-eagles/internal/dao"
	"young-eagles/internal/endpoints"
	"young-eagles/internal/services"
	"young-eagles/internal/transport"
)

func main() {
	config.LoadConfig(config.GetAppEnvLocation())

	dbConfig := dao.DbConfig{
		DbName: config.GetDbName(),
		DbUser: config.GetDbUsername(),
		DbPass: config.GetDbPass(),
		DbPort: config.GetDbPort(),
		DbHost: config.GetDbHost(),
	}

	router := mux.NewRouter().StrictSlash(true)

	v1Router := router.PathPrefix("/v1").Subrouter()

	dbConn, err := dao.InitDatabase(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.PingDb()
	if err != nil {
		log.Fatalf("error pinging error %v", err)
	}
	fmt.Printf("successfully connected to %s", dbConfig.DbName)

	pilotService := services.NewPilotService(dao.NewPilotDao(dbConn))
	pilotEndpoints := endpoints.NewPilotEndpoints(pilotService)
	transport.PostPilotData(pilotEndpoints, v1Router)
	transport.GetPilotData(pilotEndpoints, v1Router)
	transport.PatchPilotData(pilotEndpoints, v1Router)

	planesService := services.MakePlanesService(dao.NewPlaneDao(dbConn))
	planesEndpoint := endpoints.MakePlaneEndpoints(planesService)
	transport.AddPlaneDatum(planesEndpoint, v1Router)
	transport.DeletePlaneDatum(planesEndpoint, v1Router)
	transport.ReinstatePlaneDatum(planesEndpoint, v1Router)

	parentDemographicsService := services.NewParentService(dao.NewParentDao(dbConn))
	parentEndpoint := endpoints.MakeParentEndpoints(parentDemographicsService)
	transport.PostParentInformation(parentEndpoint, v1Router)

	port := fmt.Sprintf(":%s", "8080")

	srv := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting http server %v", err)
	}
}
