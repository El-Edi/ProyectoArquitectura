package routes

import (
	"arquitectura/proyectoArquitectura/controllers"

	"github.com/gin-gonic/gin"
)

func UsuarioRouter(router *gin.Engine) {

	routes := router.Group("api/v1/usuario")
	routes.POST("", controllers.UsuarioCreate)
	routes.GET("", controllers.UsuarioGet)
}
