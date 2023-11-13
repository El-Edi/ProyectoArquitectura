package routes

import (
	"arquitectura/proyectoArquitectura/controllers"

	"github.com/gin-gonic/gin"
)

func UbicacionRouter(router *gin.Engine) {

	routes := router.Group("api/v1/ubicacion")
	routes.GET("", controllers.UbicacionGet)
	routes.GET("/:id", controllers.UbicacionGetById)
}
