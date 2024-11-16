package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie


func getmovies(w http.ResponseWriter, r http.Request){
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(movies)

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		Id:    "1",
		Isbn:  "5625878",
		Title: "Movie 1",
		Director: &Director{
			Firstname: "Hosanna",
			Lastname:  "Aime",
		},
	})
	movies = append(movies, Movie{
		Id:    "2",
		Isbn:  "5645878",
		Title: "Movie 2",
		Director: &Director{
			Firstname: "Sugira",
			Lastname:  "Herve",
		},
	})

	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getmovie).Methods("GET")
	r.HandleFunc("/movies/", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")

	fmt.Println("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
