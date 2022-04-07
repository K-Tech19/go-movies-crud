package main

import (
	"encoding/json" // package to help encode our file to json when using thunder client
	"fmt"           // package to allow us to print strings and other things to the user
	"log"           // package to log
	"math/rand"
	"strconv"

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

// Database where all the movies are kept
var movies []Movie 

// Create all movies function to "get" all movies

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json") // setting our content to display as json
	json.NewEncoder(w).Encode(movies) // we are going encode w into json then we want to pass the complete movies 
}
 // Create the delete function to "delete" one movie

func deleteMovie( w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	params :=  mux.Vars(r) // params or the id

	// for range function we have to pass it an index or var i = 0 and the item in this case a movie
	for index, item := range movies { // similar to for each in other languages
			// item.ID selects a single id out of all the ids in the movies list
		if item.ID == params["id"] {

			// the movie you would want to remove gets in the first param of append
			// then we are moving the rest of the movies in the array forward by one

			movies = append(movies[:index], movies[index+1:]...) 
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // returns all the rest of the movies after one was deleted successfully 
}

// Create the getMovie function to "get" one movie

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// if you dont want to use anything as an index you an replace the keyword index with _ instead.
	for _, item := range movies{
		if item.ID == params["id"]{
			// gets one particular item or movies at a time and return to continue the task
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // decode json r = request body. Decode the movie struct
	movie.ID = strconv.Itoa(rand.Intn(1000000)) // creates and randomizes the id and converts into a string the number 100000 and 0
	movies = append(movies, movie) // adds the movie to the other movies
	json.NewEncoder(w).Encode(movie) // returns movies after we finish altering it
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	// Set json content type 
	// params
	// loop over our movies
	// delete the movie with the id that we sent
	// add a new movie - the movie that we send in the body of postman

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies{

		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie // want to append the current movie
			// below is the same idea we are doing in create movie func
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"] // different from the createMovie func, movie.ID is going keep the same id we are changing
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie) // want to return the movie to the user
			return
		}

	}
}




func main(){
	r := mux.NewRouter() // can be found inside the gorilla mux package, creates a new router to use

	movies = append(movies, Movie{ ID: "1", Isbn: "435229", Title: "Astro-Kid", Director: &Director{Firstname: "Kenny", Lastname: "Jean"}})
	movies = append(movies, Movie{ID: "2", Isbn: "234392", Title: "Spaceman in Space", Director: &Director{Firstname: "Peter", Lastname: "Hawkings"}})

	// below we are connecting our handle functions for our routes: endpoints, functions, and methods.
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at 8000")
	log.Fatal(http.ListenAndServe(":8000",r))



}