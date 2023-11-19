package routes

import (
	"arquitectura/proyectoArquitectura/controllers"

	"github.com/gin-gonic/gin"
)

func EventoRouter(router *gin.Engine) {

	routes := router.Group("api/v1/evento")
	routes.POST("", controllers.UsuarioCreate)
	routes.GET("", controllers.UsuarioGet)
}