package routers
import(
	"pelis/controler"
	"github.com/gin-gonic/gin"
)

func LoadRouters()(*gin.Engine ){

	routers := gin.Default() 

	routers.GET("api/movies/", controler.GetAllMovies)
	routers.GET("api/movies/:id",controler.GetMoviesById)

	return routers
}