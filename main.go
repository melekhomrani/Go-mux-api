package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func main() {

	movies = append(movies, Movie{Id: "1", Title: "The Shawshank Redemption", Director: &Director{Firstname: "Frank", Lastname: "Darabont"}})
	movies = append(movies, Movie{Id: "2", Title: "The Godfather", Director: &Director{Firstname: "Francis", Lastname: "Ford Coppola"}})
	movies = append(movies, Movie{Id: "3", Title: "The Godfather: Part II", Director: &Director{Firstname: "Francis", Lastname: "Ford Coppola"}})
	movies = append(movies, Movie{Id: "4", Title: "The Dark Knight", Director: &Director{Firstname: "Christopher", Lastname: "Nolan"}})
	movies = append(movies, Movie{Id: "5", Title: "12 Angry Men", Director: &Director{Firstname: "Sidney", Lastname: "Lumet"}})
	movies = append(movies, Movie{Id: "6", Title: "Schindler's List", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}})

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		panic(err)
	}
}
