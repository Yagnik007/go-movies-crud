package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"` //*Director
}

type Director struct {
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	//Adding Initial Movies.
	movies = append(movies,
		Movie{ID: "1", Isbn: "1243123", Title: "Interstellar", Director: &Director{FirstName: "Chris", LastName: "Nolan"}},
		Movie{ID: "2", Isbn: "345453", Title: "Dune", Director: &Director{FirstName: "Danie", LastName: "Villenue"}})

	
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting the server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
