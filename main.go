package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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


func getmovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(movies)

}
func deletemovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func getmovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

} 
func updatemovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")

	params := mux.Vars(r)

	var updatedMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&updatedMovie)

	for index, item := range movies {
		if item.Id == params["id"] {
			if updatedMovie.Id == "" {
				updatedMovie.Id = strconv.Itoa(rand.Intn(100))
			}
			movies[index] = updatedMovie
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
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
