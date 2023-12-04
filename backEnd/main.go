package main

import (
	"log" // Importa el paquete log de la biblioteca estándar

	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/app"
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/db"
)

func main() {
	// Iniciar la base de datos
	db.StartDbEngine()
	log.Println("Database engine started") // Cambiado a log.Println

	// Iniciar el enrutamiento y el servidor web
	app.StartRoute()
	log.Println("Application started") // Cambiado a log.Println
}
