package router

import (
	"chesscom-copy/internal/pages"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Charger les templates HTML
	r.LoadHTMLGlob("templates/*")

	r.GET("/", pages.HomePage)

	return r
}
