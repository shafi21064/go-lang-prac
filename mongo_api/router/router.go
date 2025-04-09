package router

import (
	"github.com/gorilla/mux"
	"github.com/shafi21064/go_mongo_api/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMoviesController).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovieCotroller).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatchedControler).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteSingleMovieController).Methods("DELETE")
	router.HandleFunc("/api/movies/delete", controller.DeleteAllMovieController).Methods("DELETE")
	return router
}
