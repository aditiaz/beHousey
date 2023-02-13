package routes

import (
	"housey/handlers"
	"housey/pkg/mysql"
	"housey/repositories"

	// "housey/pkg/middleware"

	"github.com/gorilla/mux"
)

func FilterRoutes(r *mux.Router) {
	filterRepository := repositories.RepositoryFilter(mysql.DB)
	h := handlers.HandlerFilter(filterRepository)

	r.HandleFunc("/multifilter", h.MultiFilter).Methods("GET")

}
