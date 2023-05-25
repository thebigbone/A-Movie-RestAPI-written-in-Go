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

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Genre    string    `json:"genre"`
	Director *Director `json:"director"`
}

type Director struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Title: "Bird Box", Genre: "Thriller", Director: &Director{
		Name: "Chris Nollan",
		Age:  50,
	}})

	movies = append(movies, Movie{ID: "2", Title: "Parasite", Genre: "Physcology", Director: &Director{
		Name: "Jack Morris",
		Age:  45,
	}})

	r.HandleFunc("/movies", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("serving at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
