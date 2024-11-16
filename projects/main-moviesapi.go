package main

// Importing necessary packages for encoding/decoding JSON, handling HTTP requests, and logging
import (
	"encoding/json" // For encoding and decoding JSON data
	"fmt"         // For formatting strings
	"github.com/gorilla/mux" // For handling HTTP requests and routing
	"log"         // For logging errors and messages
	"net/http"    // For handling HTTP requests and responses
)

// Defining a struct to represent a Movie
type Movie struct {
	Id       string    `json:"id"` // The ID of the movie
	Isbn     string    `json:"isbn"` // The ISBN of the movie
	Title    string    `json:"title"` // The title of the movie
	Director *Director `json:"director"` // The director of the movie
}

// Defining a struct to represent a Director
type Director struct {
	Firstname string `json:"firstname"` // The director's first name
	Lastname  string `json:"lastname"` // The director's last name
}

// Initializing a slice to hold all movies
var movies []Movie

// Function to handle GET request for all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	// Setting the content type of the response to JSON
	w.Header().Set("Content-Type", "application/json")
	// Encoding the movies slice to JSON and sending it in the response
	json.NewEncoder(w).Encode(movies)
}

// Function to handle GET request for a single movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	// Setting the content type of the response to JSON
	w.Header().Set("Content-Type", "application/json")
	// Extracting the ID from the URL parameters
	params := mux.Vars(r)
	// Iterating through the movies slice to find the movie with the matching ID
	for _, movie := range movies {
		if movie.Id == params["id"] {
			// If found, encoding the movie to JSON and sending it in the response
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	// If not found, encoding an empty Movie struct to JSON and sending it in the response
	json.NewEncoder(w).Encode(&Movie{})
}

// Function to handle POST request to create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	// Setting the content type of the response to JSON
	w.Header().Set("Content-Type", "application/json")
	// Decoding the JSON from the request body into a new Movie struct
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	// Appending the new movie to the movies slice
	movies = append(movies, newMovie)
	// Encoding the new movie to JSON and sending it in the response
	json.NewEncoder(w).Encode(newMovie)
}

// Function to handle PUT request to update a movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Setting the content type of the response to JSON
	w.Header().Set("Content-Type", "application/json")
	// Extracting the ID from the URL parameters
	params := mux.Vars(r)
	// Iterating through the movies slice to find the movie with the matching ID
	for index, movie := range movies {
		if movie.Id == params["id"] {
			// If found, decoding the JSON from the request body into the movie at the found index
			_ = json.NewDecoder(r.Body).Decode(&movies[index])
			// Encoding the updated movie to JSON and sending it in the response
			json.NewEncoder(w).Encode(movies[index])
			return
		}
	}
	// If not found, encoding an empty Movie struct to JSON and sending it in the response
	json.NewEncoder(w).Encode(&Movie{})
}

// Function to handle DELETE request to delete a movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Setting the content type of the response to JSON
	w.Header().Set("Content-Type", "application/json")
	// Extracting the ID from the URL parameters
	params := mux.Vars(r)
	// Iterating through the movies slice to find the movie with the matching ID
	for index, movie := range movies {
		if movie.Id == params["id"] {
			// If found, removing the movie from the slice
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// Encoding the updated movies slice to JSON and sending it in the response
	json.NewEncoder(w).Encode(movies)
}

func main() {
	// Creating a new router
	r := mux.NewRouter()

	// Initializing the movies slice with some sample data
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

	// Registering routes for handling different HTTP requests
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("starting server at port 8000\n")
	// Starting the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
