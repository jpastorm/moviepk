// genre.model.go
package model

type Movie_Genre struct {
	Movie_Genres int `json:"mov_genres" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Gen_id       int `json:"gen_id"`
	Mov_id       int `json:"mov_id"`
}
