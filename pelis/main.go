package main

import (
	"fmt"
	"pelis/routers"
)


func main(){
	fmt.Println("Iniciado")
	run := routers.LoadRouters()
	run.Run("localhost:3006")
}