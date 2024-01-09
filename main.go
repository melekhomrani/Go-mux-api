package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type (
	Movie struct {
		Id       string    `json:"id"`
		Title    string    `json:"title"`
		Director *Director `json:"director"`
	}

	Director struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}
)

var movies []Movie
var directors []Director

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	lastId, _ := strconv.Atoi(movies[len(movies)-1].Id)
	movie.Id = strconv.Itoa(lastId + 1)
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	for index, item := range movies {
		if item.Id == params["id"] {
			movie.Id = params["id"]
			movies[index] = movie
			break
		}
	}
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode("Movie with id " + params["id"] + " has been deleted successfully")
}

func getDirectorByMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item.Director)
			return
		}
	}
}

func main() {

	directors = append(directors, Director{Firstname: "Frank", Lastname: "Darabont"})
	directors = append(directors, Director{Firstname: "Francis", Lastname: "Ford Coppola"})
	directors = append(directors, Director{Firstname: "Christopher", Lastname: "Nolan"})
	directors = append(directors, Director{Firstname: "Sidney", Lastname: "Lumet"})
	directors = append(directors, Director{Firstname: "Steven", Lastname: "Spielberg"})

	movies = append(movies, Movie{Id: "1", Title: "The Shawshank Redemption", Director: &Director{Firstname: "Frank", Lastname: "Darabont"}})
	movies = append(movies, Movie{Id: "2", Title: "The Godfather", Director: &Director{Firstname: "Francis", Lastname: "Ford Coppola"}})
	movies = append(movies, Movie{Id: "3", Title: "The Godfather: Part II", Director: &Director{Firstname: "Francis", Lastname: "Ford Coppola"}})
	movies = append(movies, Movie{Id: "4", Title: "The Dark Knight", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{Id: "5", Title: "12 Angry Men", Director: &Director{Firstname: "Sidney", Lastname: "Lumet"}})
	movies = append(movies, Movie{Id: "6", Title: "Schindler's List", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", addMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}/directors", getDirectorByMovie).Methods("GET")

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		panic(err)
	}
}
