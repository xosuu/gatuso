package controler

import (
	"net/http"
	domain "pelis/domain/movie"
	"strconv"

	"github.com/gin-gonic/gin"
)


var movies = []domain.Movie{
	{Id: 1, Name: "rocky"},
	{Id: 2, Name: "It"},
	{Id: 3, Name: "Transformers"},
	{Id: 4, Name: "Pele"},
}

//get All movies
func GetAllMovies(c *gin.Context){

	c.IndentedJSON(http.StatusOK, movies)
}


//get Movies by id 
func GetMoviesById(c *gin.Context){
	id, _:= strconv.Atoi(c.Param("id"))

	for _, v := range(movies){
		if( id == v.Id){
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	
}









