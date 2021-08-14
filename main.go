package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"rest-api-mysql/controller"
	"rest-api-mysql/database"
)

func main() {
	config := database.Config{
		ServerName: "localhost:3306",
		User:       "admin",
		Password:   "admin",
		DB:         "dbmusic",
	}
	connectionString := database.GetConnectionString(config)
	database.Connect(connectionString)
	defer database.Db.Close()

	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/artists", controller.GetArtists).Methods("GET")

	router.HandleFunc("/artist", controller.CreateArtist).Methods("POST")
	router.HandleFunc("/artist/{id}", controller.GetArtist).Methods("GET")
	router.HandleFunc("/artist/{id}", controller.UpdateArtist).Methods("PUT")
	router.HandleFunc("/artist/{id}", controller.DeleteArtist).Methods("DELETE")

	http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(router))

}
