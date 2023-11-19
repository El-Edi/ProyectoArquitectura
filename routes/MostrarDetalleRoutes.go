package routes

import (
	"arquitectura/proyectoArquitectura/controllers"

	"github.com/gin-gonic/gin"
)

func MostrarDetalleRouter(router *gin.Engine) {

	routes := router.Group("api/v1/detalle")
	routes.GET("/:id", controllers.DetalleEventoGetById)
}
