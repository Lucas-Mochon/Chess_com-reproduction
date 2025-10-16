package main

import (
	"chesscom-copy/internal/router"
	"log"
)

func main() {
	r := router.SetupRouter()
	log.Println("Serveur démarré sur http://localhost:8080")
	r.Run(":8080")
}
