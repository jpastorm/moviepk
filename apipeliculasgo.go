package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	conn "github.com/moviepk/config"
	c "github.com/moviepk/controllers"
	m "github.com/moviepk/models"
)

var err error

func main() {

	conn.Db, err = gorm.Open("mysql", "root:password@tcp(127.0.0.1:3306)/apigopeliculas?charset=utf8&parseTime=True")
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function

	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}
	conn.Db.AutoMigrate(&m.Genres{})
	conn.Db.AutoMigrate(&m.Movie{})
	conn.Db.AutoMigrate(&m.Movie_Genre{})
	conn.Db.Table("movie_genres").AddForeignKey("gen_id", "genres(gen_id)", "RESTRICT", "RESTRICT")
	conn.Db.Table("movie_genres").AddForeignKey("mov_id", "movies(mov_id)", "RESTRICT", "RESTRICT")

	myRouter := mux.NewRouter().StrictSlash(true)
	log.Println("Starting development server at http://127.0.0.1:3000/")
	log.Println("Quit the server with CONTROL-C.")

	headers := handlers.AllowedHeaders([]string{"x-Requested-with", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	myRouter.HandleFunc("/genres", c.CreateNewGenre).Methods("POST")
	myRouter.HandleFunc("/genres", c.ReturnAllGenres).Methods("GET")
	myRouter.HandleFunc("/genres/{id}", c.ReturnSinbleGenresdos).Methods("GET")
	myRouter.HandleFunc("/genres/{id}", c.DeleteGenre).Methods("DELETE")
	myRouter.HandleFunc("/genres/{id}", c.UpdateGenre).Methods("PUT")

	myRouter.HandleFunc("/movies", c.CreateNewMovies).Methods("POST")
	myRouter.HandleFunc("/movies", c.ReturnAllMovies).Methods("GET")
	myRouter.HandleFunc("/movies/{id}", c.ReturnSinbleMoviesdos).Methods("GET")
	myRouter.HandleFunc("/movies/{id}", c.DeleteMovies).Methods("DELETE")
	myRouter.HandleFunc("/movies/{id}", c.UpdateMovies).Methods("PUT")

	myRouter.HandleFunc("/moviegenre", c.CreateNewMovieGenre).Methods("POST")
	myRouter.HandleFunc("/moviegenre", c.ReturnAllMovieGenre).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(myRouter)))
}
