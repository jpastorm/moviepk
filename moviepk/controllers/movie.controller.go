// movie.controller.go
package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	conn "github.com/moviepk/config"
	m "github.com/moviepk/models"
)

//var db *gorm.DB

//var Db = db
/*
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}*/
func CreateNewMovies(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(reqBody)
	var movie m.Movie
	json.Unmarshal(reqBody, &movie)
	conn.Db.Create(&movie)
	fmt.Println("Creating a new Movie")
	json.NewEncoder(w).Encode(movie)
}
func ReturnAllMovies(w http.ResponseWriter, r *http.Request) {
	movie := []m.Movie{}
	conn.Db.Find(&movie)
	fmt.Println("Return all Movies")
	json.NewEncoder(w).Encode(movie)
}
func ReturnSinbleMoviesdos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	movie := []m.Movie{}
	s, err := strconv.Atoi(key)
	if err == nil {
		conn.Db.Where("mov_id = ?", s).Find(&movie)
		json.NewEncoder(w).Encode(movie)
	}

}
func ReturnSingleMovies(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	movie := []m.Movie{}
	conn.Db.Find(&movie)
	for _, movie := range movie {
		// string to int
		s, err := strconv.Atoi(key)
		if err == nil {
			if movie.Mov_id == s {
				fmt.Println(movie)
				fmt.Println("Movie", key)
				json.NewEncoder(w).Encode(movie)
			}
		}
	}
}
func DeleteMovies(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	s, err := strconv.Atoi(key)
	if err == nil {
		conn.Db.Delete(&m.Movie{Mov_id: s})
		fmt.Println("Delete")
	}

}
func UpdateMovies(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data m.Movie
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	vars := mux.Vars(r)
	key := vars["id"]
	s, er := strconv.Atoi(key)

	if er == nil {
		data.Mov_id = s
		conn.Db.Save(&data)
		fmt.Println("Actualizado", data)
	}

}
