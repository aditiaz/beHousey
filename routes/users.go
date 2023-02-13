package routes

import (
	"housey/handlers"
	"housey/pkg/middleware"
	"housey/pkg/mysql"
	"housey/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)
	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/updateuser", middleware.UploadFile(h.UpdateUser)).Methods("PATCH")
	r.HandleFunc("/changepassword", middleware.Auth(h.ChangePassword)).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")

}
