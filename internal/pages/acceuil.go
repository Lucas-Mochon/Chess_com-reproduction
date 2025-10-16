package pages

import (
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(200, "accueil.html", gin.H{
		"title": "Bienvenue sur ChessGo",
	})
}
