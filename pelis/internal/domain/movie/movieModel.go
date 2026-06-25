package movie
import (

	"gorm.io/gorm"
)




type Movie struct{
	gorm.Model
	Name string 
	MovieUrl string 
	PosterUrl string
	Duration string 
	Sinopsis string 
	Genre string 
}