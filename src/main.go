package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mndozkn/go-movies-crud/handlers"
	"github.com/mndozkn/go-movies-crud/types"
)

type Movie = types.Movie
type Director = types.Director

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438277", Title: "Movie One", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "431232", Title: "Movie Two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})
	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", handlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
