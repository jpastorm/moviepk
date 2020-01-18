// genre.controller.go
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

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}
func CreateNewGenre(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(reqBody)
	var genres m.Genres
	json.Unmarshal(reqBody, &genres)
	conn.Db.Create(&genres)
	fmt.Println("Creating a new Genre")
	json.NewEncoder(w).Encode(genres)
}
func ReturnAllGenres(w http.ResponseWriter, r *http.Request) {
	genres := []m.Genres{}
	conn.Db.Find(&genres)
	fmt.Println("Return all Genres")
	json.NewEncoder(w).Encode(genres)
}
func ReturnSinbleGenresdos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	genres := []m.Genres{}
	s, err := strconv.Atoi(key)
	if err == nil {
		conn.Db.Where("gen_id = ?", s).Find(&genres)
		json.NewEncoder(w).Encode(genres)
	}

}
func ReturnSingleGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	genres := []m.Genres{}
	conn.Db.Find(&genres)
	for _, genres := range genres {
		// string to int
		s, err := strconv.Atoi(key)
		if err == nil {
			if genres.Gen_id == s {
				fmt.Println(genres)
				fmt.Println("Genre", key)
				json.NewEncoder(w).Encode(genres)
			}
		}
	}
}
func DeleteGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	s, err := strconv.Atoi(key)
	if err == nil {
		conn.Db.Delete(&m.Genres{Gen_id: s})
		fmt.Println("Delete")
	}

}
func UpdateGenre(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data m.Genres
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	vars := mux.Vars(r)
	key := vars["id"]
	s, er := strconv.Atoi(key)

	if er == nil {
		data.Gen_id = s
		conn.Db.Save(&data)
		fmt.Println("Actualizado", data)
	}

}
