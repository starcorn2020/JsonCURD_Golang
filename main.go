package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{

	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Tittle string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)

}

func deleteMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	params := mux.Var(r)
	for index, item := range movies{

		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	params := mux.Var(r)
	for _, item := range movies{
		
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_ = json.NewEncoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type","application/json")
	params := mux.Var(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode((movie))
			return
		}
	}
}

func main(){
	
	r := mux.NewRouuter()

	movies = append(movies, Movie{ID:"1", Isbn:"43822", Tittle:"Movie One", Director: &Director{Firstname:"John",Lastname:"Doe"}})
	movies = append(movies, Movie{ID:"2", Isbn:"43822112", Tittle:"Movie Two", Director: &Director{Firstname:"Steve",Lastname:"Wu"}})

	r.HandleFunc("/movies",getMovies).Method("GET")
	r.HandleFunc("/movies(id)",getMovie).Method("GET")
	r.HandleFunc("/movies",createMovie).Method("POST")
	r.HandleFunc("/movies(id)",updateMovie).Method("PUT")
	r.HandleFunc("/movies(id)",deleteMovie).Method("DELETE")

	fmt.Printf("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080",r))
}