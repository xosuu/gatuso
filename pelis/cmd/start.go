package main

import (
	"fmt"
	"net/http"
	"pelis/internal/config"
	"pelis/internal/controler"
	"pelis/internal/domain/movie"
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB
func Init(){

	port := ":3006"
	DB = config.LoadDatabase()
	Repo := &movie.MovieRepository{Db: DB}
	Controller := &controler.MovieController{Repo: *Repo}
	router := config.LoadRouters(Controller)

	DB.AutoMigrate(&movie.Movie{})


	server := &http.Server{
		Addr: port,
		Handler: router,
		ReadHeaderTimeout: 2 * time.Second,
	}

	err := server.ListenAndServe()
	if( err != nil){
		fmt.Println(err)
		return
	}

	fmt.Println("Server iniciado")

	
}