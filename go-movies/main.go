package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

// Get Movies - ALL
func getAllMovies(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, true, "Movies retrieved successfully", movies)
}

// For Getting a Single Movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			respondJSON(w, http.StatusOK, true, "Movie retrieved successfully", item)
			return
		}
	}
	respondJSON(w, http.StatusNotFound, false, "Movie not found", nil)
}

// For Deleting a Movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			respondJSON(w, http.StatusOK, true, "Movie deleted successfully", movies)
			return
		}
	}
	respondJSON(w, http.StatusNotFound, false, "Movie not found", nil)
}

// For Creating a Movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondJSON(w, http.StatusBadRequest, false, "Invalid request payload", nil)
		return
	}

	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	respondJSON(w, http.StatusCreated, true, "Movie created successfully", movie)
}

// For Updating a Movie
func updateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Find the movie to update
	for index, item := range movies {
		if item.ID == params["id"] {
			// Remove the old movie
			movies = append(movies[:index], movies[index+1:]...)

			// Create updated movie
			var movie Movie
			if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
				respondJSON(w, http.StatusBadRequest, false, "Invalid request payload", nil)
				return
			}

			movie.ID = params["id"]
			movies = append(movies, movie)
			respondJSON(w, http.StatusOK, true, "Movie updated successfully", movie)
			return
		}
	}
	respondJSON(w, http.StatusNotFound, false, "Movie not found", nil)
}

func respondJSON(w http.ResponseWriter, status int, success bool, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ApiResponse{
		Success: success,
		Message: message,
		Data:    data,
	})
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "12345",
		Title: "First Movie",
		Director: &Director{
			FirstName: "John",
			LastName:  "Mike",
		},
	})

	// Second Movie
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "42345",
		Title: "Second Movie",
		Director: &Director{
			FirstName: "Mike",
			LastName:  "Tyson",
		},
	})

	// Third Movie
	movies = append(movies, Movie{
		ID:    "3",
		Isbn:  "32345",
		Title: "Third Movie",
		Director: &Director{
			FirstName: "Kalesh",
			LastName:  "Damodar",
		},
	})

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting the Server at Port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
