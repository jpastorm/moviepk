// genre.model.go
package model

type Movie struct {
	Mov_id    int    `json:"mov_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Mov_title string `json:"mov_title"`
	Mov_year  string `json:"mov_year"`
}
