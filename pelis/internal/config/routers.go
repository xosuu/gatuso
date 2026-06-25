package config

import (
	"pelis/internal/controler"

	"github.com/gin-gonic/gin"
)

func LoadRouters(controller *controler.MovieController)(*gin.Engine ){

	routers := gin.Default() 

	routers.GET("api/movies/", controller.GetAllMovies)
	routers.GET("api/movies/:id", controller.GetById)
	routers.POST("api/movies", controller.InsertMovie)

	return routers
}