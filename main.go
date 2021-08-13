package main

import (
	"log"

	"github.com/ciarmail/twitter/bd"
	"github.com/ciarmail/twitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
