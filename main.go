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
	config.LoadConfig()

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

	fmt.Println("setting up endpoints")
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

	childrenService := services.NewChildrenService(dao.NewChildrenDao(dbConn))
	childrenEndpoints := endpoints.NewChildEndpoints(childrenService)
	transport.PostChildInformation(childrenEndpoints, v1Router)
	transport.GetChildInformation(childrenEndpoints, v1Router)

	port := fmt.Sprintf(":%s", "8080")

	fmt.Println("setting up server on port " + port)

	srv := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server on " + srv.Addr)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error starting http server %v", err)
	}
}
