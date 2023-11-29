package routes

import (
	"github.com/gorilla/mux"
	"github.com/iamtonmoy0/movie-crud-api/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/movies", controllers.GetAllMyMovies).Methods("GET")
	router.HandleFunc("/create", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", controllers.MarkAsWatched).Methods("POST")
	router.HandleFunc("/movie/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/movie/delete", controllers.DeleteAllMovies).Methods("DELETE")
	return router
}
