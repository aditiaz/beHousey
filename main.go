package main

import (
	"fmt"
	"housey/database"
	"housey/pkg/mysql"
	"os"

	"housey/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// godotenv
	godotenv.Load()

	// Database
	mysql.DatabaseInit()

	// Migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	// Initialization "uploads" folder to public here ...
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// Setup allowed Header, Method, and Origin for CORS on this below code ...
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})
	port := os.Getenv("PORT")
	// Embed the setup allowed in 2 parameter on this below code ...
	fmt.Println("Server is running on 5000")
	// http.ListenAndServe("localhost:5000", handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
	http.ListenAndServe(":"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))

	// handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
