package main

import (
	"net/http"

	"arquitectura/proyectoArquitectura/configs"
	"arquitectura/proyectoArquitectura/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	r := gin.Default()

	routes.UsuarioRouter(r)
	routes.UbicacionRouter(r)
	routes.MostrarDetalleRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run()
}
