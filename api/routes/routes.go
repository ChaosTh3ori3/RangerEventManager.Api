package routes

import (
	"net/http"

	"github.com/ChaosTh3ori3/RangerEventManager.Api/api/handler/camp"
	"github.com/ChaosTh3ori3/RangerEventManager.Api/config"

	"github.com/gorilla/mux"
)

type Routes struct {
	createCampHandler *camp.CreateCampHandler
	getCampHandler    *camp.GetCampHandler
	keyCloakSettings  *config.KeyCloakSettings
}

func NewRoutes(
	createCampHandler *camp.CreateCampHandler,
	getCampHandler *camp.GetCampHandler,
	keyCloakSettings *config.KeyCloakSettings) Routes {
	return Routes{
		createCampHandler: createCampHandler,
		getCampHandler:    getCampHandler,
		keyCloakSettings:  keyCloakSettings,
	}
}

func (r Routes) SetupRoutes() http.Handler {
	router := mux.NewRouter()

	// authMiddleware := middleware.NewAuthMiddleware(r.keyCloakSettings)

	// router.Use(authMiddleware.Authenticate)

	router.HandleFunc("/api/camp/{campNumber}", r.getCampHandler.GetCampByCampNumber).Methods("Get")

	router.HandleFunc("/api/plan-camp", r.createCampHandler.HandleCreateCamp).Methods("Post")

	return router
}
