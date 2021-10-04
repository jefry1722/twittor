package main

import (
	"log"

	"github.com/jefry1722/twittor/bd"
	"github.com/jefry1722/twittor/handlers"
)

func main() {
	if !bd.ChequeoConexion(){
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}