// genre.model.go
package model

type Genres struct {
	Gen_id    int    `json:"gen_id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Gen_title string `json:"gen_title"`
}
