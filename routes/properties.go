package routes

import (
	"housey/handlers"
	"housey/pkg/middleware"
	"housey/pkg/mysql"
	"housey/repositories"

	"github.com/gorilla/mux"
	// "gorm.io/gorm/migrator"
)

func PropertyRoutes( r *mux.Router) {
	propertyRepository := repositories.RepositoryProperty(mysql.DB)
	h := handlers.HandlerProperty(propertyRepository)

	r.HandleFunc("/properties",h.FindProperties).Methods("GET")
	r.HandleFunc("/property/{id}",h.GetProperty).Methods("GET")
	r.HandleFunc("/addproperty",middleware.Auth( middleware.UploadFile(h.AddProperty))).Methods("POST")
}