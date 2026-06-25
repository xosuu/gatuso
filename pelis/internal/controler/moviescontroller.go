package controler

import (
	"fmt"
	"net/http"
	"strconv"

	//"strconv"
	//"pelis/internal/config"

	"pelis/internal/domain/movie"

	"github.com/gin-gonic/gin"
)





type MovieController struct{
	Repo movie.MovieRepository
}

type obj struct{
	Message string
}

//GET BY ID
func (m *MovieController) GetById(c *gin.Context){
	id := c.Param("id")
	idd, _ := strconv.Atoi(id)
	movie, err := m.Repo.GetById(idd)
	if(err != nil){
		c.IndentedJSON(http.StatusNotFound, &obj{Message: "No existe"})
		return
	}
	c.IndentedJSON(http.StatusOK, movie)
}



//GET ALL ITEMS
func (m *MovieController) GetAllMovies(c *gin.Context){
	movies, err := m.Repo.GetAllMovies()
	if(err != nil){
		c.IndentedJSON(http.StatusNotFound, &obj{Message: "No existe papu"})
	}
	var Movies []*movie.MovieResponse
	for _, v := range movies{
		Movies = append(Movies, &movie.MovieResponse{
			Name: v.Name, MovieUrl: v.MovieUrl, PosterUrl: v.PosterUrl,
			Duration: v.Duration, Sinopsis: v.Sinopsis, Genre: v.Genre,
			
		})
	}
	fmt.Println(movies)
	c.IndentedJSON(http.StatusOK, &Movies)
}

//POST INSERT MOVIE
func (m *MovieController)InsertMovie(c *gin.Context){
	
	var MoviePost movie.MoviePost
	data := c.BindJSON(&MoviePost)
	if(data != nil){
		fmt.Println(data)
		c.IndentedJSON(http.StatusNotFound, &obj{Message: "Error jijo"})
		return
	}

	movie := &movie.Movie{Name: MoviePost.Name, MovieUrl: MoviePost.MovieUrl}
	err := m.Repo.Save(movie)
	if (err != nil){
		c.IndentedJSON(http.StatusNotFound, &obj{Message: "Error jijo"})
		return
	}
	c.IndentedJSON(http.StatusCreated, &MoviePost)
	
}













// //get Movies by id 
// func (m *MovieController) GetMoviesById(c *gin.Context){
// 	id, _:= strconv.Atoi(c.Param("id"))

// 	for _, v := range(movies){
// 		if( id == v.Id){
// 			c.IndentedJSON(http.StatusOK, v)
// 			return
// 		}
// 	}
	
// }









