package api

import (
	"caas-aks-api-go/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/clusters", handlers.CreateCluster).Methods("POST")
	r.HandleFunc("/clusters/{name}", handlers.GetCluster).Methods("GET")
	r.HandleFunc("/clusters", handlers.ListClusters).Methods("GET")
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	return r
}
