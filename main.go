package main

import (
	"log"

	"github.com/RubenStark/twitter/bd"
	"github.com/RubenStark/twitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == false {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
