// genre.controller.go
package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"strconv"

	//"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	conn "github.com/moviepk/config"
	m "github.com/moviepk/models"
)

//var db *gorm.DB

//var Db = db

func CreateNewMovieGenre(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Println(reqBody)
	var moviegenre m.Movie_Genre
	json.Unmarshal(reqBody, &moviegenre)
	conn.Db.Create(&moviegenre)
	fmt.Println("Creating a new MovieGenre")
	json.NewEncoder(w).Encode(moviegenre)
}
func ReturnAllMovieGenre(w http.ResponseWriter, r *http.Request) {
	moviegenre := []m.Movie_Genre{}
	conn.Db.Find(&moviegenre)
	fmt.Println("Return all MovieGenre")
	json.NewEncoder(w).Encode(moviegenre)
}
