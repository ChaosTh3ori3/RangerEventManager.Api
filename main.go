package main

import (
	"fmt"
	"net/http"

	"github.com/ChaosTh3ori3/RangerEventManager.Api/api/handler/camp"
	"github.com/ChaosTh3ori3/RangerEventManager.Api/api/routes"
	"github.com/ChaosTh3ori3/RangerEventManager.Api/config"
	"github.com/ChaosTh3ori3/RangerEventManager.Api/persistens/database"
	"github.com/ChaosTh3ori3/RangerEventManager.Api/persistens/repositories"
)

func main() {

	configFilePath := "config/config.yml"

	config, configErr := config.LoadConfig(configFilePath)
	if configErr != nil {
		fmt.Printf("Fehler beim Laden der Konfiguration: %s\n", configErr)
		return
	}

	database := database.NewDatabase(&config.MongoDBConnection.ConnectionString)

	database.InitDatabase()

	campRepository := repositories.NewCampRepository(database.GetCampsCollection())

	getCampHandler := camp.NewGetCampHandler(&campRepository)
	createCampHandler := camp.NewCreateCampHandler(&campRepository)

	routes := routes.NewRoutes(
		&createCampHandler,
		&getCampHandler,
		&config.KeyCloakSettings)

	router := routes.SetupRoutes()

	// Starten des HTTP-Servers auf Port 5555
	httpErr := http.ListenAndServe(":5555", router)
	if httpErr != nil {
		panic(httpErr)
	}
}
