package main

import (
	"encoding/json" // package to help encode our file to json when using thunder client
	"fmt"           // package to allow us to print strings and other things to the user
	"log"           // package to log
	// package to allow us to create randomized integers
	"net/http" // package to allow us to create a server through golang
	// package for converting int to strings

	"github.com/gorilla/mux" // an external package that we installed
)

// Structs are a way to structure and use data. It allows us to group data.
// To use a struct we declare the type of struct we are going to use.
// The code below shows how to define a struct type using the 'type' keyword.
// The '*' points to the director class and '&' references that director class for use
// golang's version of classes

//create multiple movies of type Movie
type Movie struct{ 
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"` // pointing to the director struct below
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie 

// Create all movies function

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json") // setting our content to display as json
	json.NewEncoder(w).Encode(movies) // we are going encode w into json then we want to pass the complete movies 
}
 // create the delete function

func deleteMovie( w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	params :=  mux.Vars(r) // params or the id

	// for a range function we have to pass it an index or var i = 0 and the item in this case a movie
	for index, item := range movies { // similar to for each in other languages
			// item.ID selects a single id out of all the ids in the movies list
		if item.ID == params["id"] {

			// the movie you would want to remove gets in the first param of append
			// then we are moving the rest of the movies in the array forward by one

			movies = append(movies[:index], movies[index+1:]...) 
			break
		}
	}
}



func main(){
	r := mux.NewRouter() // can be found inside the gorilla mux package, creates a new router to use

	movies = append(movies, Movie{ ID: "1", Isbn: "435229", Title: "Astro-Kid", Director: &Director{Firstname: "Kenny", Lastname: "Jean"}})
	movies = append(movies, Movie{ID: "2", Isbn: "234392", Title: "Spaceman in Space", Director: &Director{Firstname: "Peter", Lastname: "Hawkings"}})

	// below we are creating our handle functions for our routes, methods, and functions
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at 8000")
	log.Fatal(http.ListenAndServe(":8000",r))



}